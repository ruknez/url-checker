package checker

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	entity "url-checker/internal/domain"
	"url-checker/internal/service/checker/mocks"
)

func Test_checker_GetStatus(t *testing.T) {
	tmpTime := time.Now()

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
		wantErr assert.ErrorAssertionFunc
		asserts func(t *testing.T, args args, fields fields)
	}{
		{
			name: "success",
			fields: fields{
				urlRepo: &mocks.UrlRepositoryMock{
					GetFunc: func(ctx context.Context, url string) (*entity.UrlInfo, error) {
						assert.Equal(t, "http://google.com", url)
						return &entity.UrlInfo{
							URL:       "url",
							Duration:  10 * time.Second,
							LastCheck: &tmpTime,
							Status:    entity.Available,
						}, nil
					},
				},
				tickDuration:    time.Second,
				logger:          &mocks.LoggerMock{},
				statuserService: &mocks.GetUrlStatuserMock{},
			},
			args: args{
				ctx: context.Background(),
				url: "http://google.com",
			},
			want:    entity.Available,
			wantErr: assert.NoError,
			asserts: func(t *testing.T, args args, fields fields) {
				assert.Equal(t, len(fields.urlRepo.GetCalls()), 1)
			},
		},
		{
			name: "success",
			fields: fields{
				urlRepo: &mocks.UrlRepositoryMock{
					GetFunc: func(ctx context.Context, url string) (*entity.UrlInfo, error) {
						assert.Equal(t, "http://google.com", url)
						return &entity.UrlInfo{
							URL:       "url",
							Duration:  10 * time.Second,
							LastCheck: &tmpTime,
							Status:    entity.Available,
						}, errors.New("some error")
					},
				},
				tickDuration:    time.Second,
				logger:          &mocks.LoggerMock{},
				statuserService: &mocks.GetUrlStatuserMock{},
			},
			args: args{
				ctx: context.Background(),
				url: "http://google.com",
			},
			want: entity.NotCheck,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "some error")
			},
			asserts: func(t *testing.T, args args, fields fields) {
				assert.Equal(t, len(fields.urlRepo.GetCalls()), 1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &checker{
				urlRepo:         tt.fields.urlRepo,
				tickDuration:    tt.fields.tickDuration,
				logger:          tt.fields.logger,
				statuserService: tt.fields.statuserService,
			}
			got, err := c.GetStatus(tt.args.ctx, tt.args.url)
			tt.asserts(t, tt.args, tt.fields)

			if !tt.wantErr(t, err, fmt.Sprintf("GetStatus(%v, %v)", tt.args.ctx, tt.args.url)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetStatus(%v, %v)", tt.args.ctx, tt.args.url)
		})
	}
}
