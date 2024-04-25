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
