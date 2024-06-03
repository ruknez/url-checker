package checker

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	entity "url-checker/internal/domain"
	"url-checker/internal/service/checker/mocks"
)

func Test_checker_checkUrl(t *testing.T) {
	type fields struct {
		urlRepo         *mocks.UrlRepositoryMock
		tickDuration    time.Duration
		logger          *mocks.LoggerMock
		statuserService *mocks.GetUrlStatuserMock
	}
	type args struct {
		ctx context.Context
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Status
		wantErr string
		asserts func(t *testing.T, args args, fields fields)
	}{
		{
			name: "success",
			fields: fields{
				urlRepo:      &mocks.UrlRepositoryMock{},
				tickDuration: time.Second,
				logger:       &mocks.LoggerMock{},
				statuserService: &mocks.GetUrlStatuserMock{
					GetUrlStatusFunc: func(ctx context.Context, url string) (entity.Status, error) {
						assert.Equal(t, "https://google.com", url)

						return entity.Available, nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
				url: "https://google.com",
			},
			want:    entity.Available,
			wantErr: "",
			asserts: func(t *testing.T, args args, fields fields) {
				assert.Equal(t, len(fields.statuserService.GetUrlStatusCalls()), 1)
			},
		},
		{
			name: "error",
			fields: fields{
				urlRepo:      &mocks.UrlRepositoryMock{},
				tickDuration: time.Second,
				logger:       &mocks.LoggerMock{},
				statuserService: &mocks.GetUrlStatuserMock{
					GetUrlStatusFunc: func(ctx context.Context, url string) (entity.Status, error) {
						assert.Equal(t, "https://google.com", url)

						return entity.Available, errors.New("someError")
					},
				},
			},
			args: args{
				ctx: context.Background(),
				url: "https://google.com",
			},
			want:    entity.Available,
			wantErr: "statuserService.GetUrlStatus",
			asserts: func(t *testing.T, args args, fields fields) {
				assert.Equal(t, len(fields.statuserService.GetUrlStatusCalls()), 1)

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewChecker(
				context.Background(),
				tt.fields.urlRepo,
				tt.fields.logger,
				tt.fields.tickDuration,
				tt.fields.statuserService,
			)
			got, err := c.checkUrl(tt.args.ctx, tt.args.url)

			if tt.wantErr != "" {
				assert.ErrorContains(t, err, tt.wantErr)
			} else {
				assert.Equal(t, tt.want, got)
				assert.NoError(t, err)
			}

			tt.asserts(t, tt.args, tt.fields)
		})
	}
}

func TestChecker_checkAllUrls(t *testing.T) {
	type fields struct {
		urlRepo         *mocks.UrlRepositoryMock
		tickDuration    time.Duration
		logger          *mocks.LoggerMock
		statuserService *mocks.GetUrlStatuserMock
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		asserts func(t *testing.T, args args, fields fields)
	}{
		{
			name: "success",
			fields: fields{
				urlRepo: &mocks.UrlRepositoryMock{
					GetAllUrlsFunc: func(ctx context.Context) []string {
						return []string{"url_1"}
					},
					UpdateStatusFunc: func(ctx context.Context, url string, status entity.Status) error {
						return nil
					},
				},
				tickDuration: time.Second * 1,
				logger:       &mocks.LoggerMock{},
				statuserService: &mocks.GetUrlStatuserMock{
					GetUrlStatusFunc: func(ctx context.Context, url string) (entity.Status, error) {
						return entity.Available, nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			asserts: func(t *testing.T, args args, fields fields) {
				assert.Equal(t, len(fields.urlRepo.GetAllUrlsCalls()), 1)
				assert.Equal(t, len(fields.urlRepo.UpdateStatusCalls()), 1)
				assert.Equal(t, len(fields.statuserService.GetUrlStatusCalls()), 1)

				assert.Equal(t, fields.urlRepo.UpdateStatusCalls()[0].URL, "url_1")
				assert.Equal(t, fields.urlRepo.UpdateStatusCalls()[0].Status, entity.Available)

				assert.Equal(t, fields.statuserService.GetUrlStatusCalls()[0].URL, "url_1")
			},
		},
		{
			name: "checkUrl error",
			fields: fields{
				urlRepo: &mocks.UrlRepositoryMock{
					GetAllUrlsFunc: func(ctx context.Context) []string {
						return []string{"url_1"}
					},
					UpdateStatusFunc: func(ctx context.Context, url string, status entity.Status) error {
						return nil
					},
				},
				tickDuration: time.Second * 1,
				logger: &mocks.LoggerMock{
					ErrorFunc: func(ctx context.Context, args ...interface{}) {
					},
				},
				statuserService: &mocks.GetUrlStatuserMock{
					GetUrlStatusFunc: func(ctx context.Context, url string) (entity.Status, error) {
						return entity.Available, errors.New("checkUrlError")
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			asserts: func(t *testing.T, args args, fields fields) {
				assert.Equal(t, len(fields.urlRepo.GetAllUrlsCalls()), 1)
				assert.Equal(t, len(fields.statuserService.GetUrlStatusCalls()), 1)
				assert.Equal(t, len(fields.urlRepo.UpdateStatusCalls()), 0)

				assert.Equal(t, fields.statuserService.GetUrlStatusCalls()[0].URL, "url_1")

				assert.Equal(t, len(fields.logger.ErrorCalls()), 1)

				assert.ErrorContains(t, fields.logger.ErrorCalls()[0].Args[0].(error), "checkUrlError")
				assert.ErrorContains(t, fields.logger.ErrorCalls()[0].Args[0].(error), "url_1")
			},
		},
		{
			name: "UpdateStatus error",
			fields: fields{
				urlRepo: &mocks.UrlRepositoryMock{
					GetAllUrlsFunc: func(ctx context.Context) []string {
						return []string{"url_1"}
					},
					UpdateStatusFunc: func(ctx context.Context, url string, status entity.Status) error {
						return errors.New("UpdateStatusError")
					},
				},
				tickDuration: time.Second * 1,
				logger: &mocks.LoggerMock{
					ErrorFunc: func(ctx context.Context, args ...interface{}) {
					},
				},
				statuserService: &mocks.GetUrlStatuserMock{
					GetUrlStatusFunc: func(ctx context.Context, url string) (entity.Status, error) {
						return entity.Available, nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			asserts: func(t *testing.T, args args, fields fields) {
				assert.Equal(t, len(fields.urlRepo.GetAllUrlsCalls()), 1)
				assert.Equal(t, len(fields.statuserService.GetUrlStatusCalls()), 1)
				assert.Equal(t, len(fields.urlRepo.UpdateStatusCalls()), 1)

				assert.Equal(t, fields.urlRepo.UpdateStatusCalls()[0].URL, "url_1")
				assert.Equal(t, fields.urlRepo.UpdateStatusCalls()[0].Status, entity.Available)
				assert.Equal(t, fields.statuserService.GetUrlStatusCalls()[0].URL, "url_1")

				assert.Equal(t, len(fields.logger.ErrorCalls()), 1)

				assert.ErrorContains(t, fields.logger.ErrorCalls()[0].Args[0].(error), "UpdateStatusError")
				assert.ErrorContains(t, fields.logger.ErrorCalls()[0].Args[0].(error), "url_1")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Checker{
				urlRepo:         tt.fields.urlRepo,
				tickDuration:    tt.fields.tickDuration,
				logger:          tt.fields.logger,
				statuserService: tt.fields.statuserService,
			}
			c.checkAllUrls(tt.args.ctx)
			tt.asserts(t, tt.args, tt.fields)
		})
	}
}
