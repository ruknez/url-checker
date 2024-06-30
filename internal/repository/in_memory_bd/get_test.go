package in_memory_bd

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	entity "url-checker/internal/domain"
	inMemoryBd "url-checker/internal/repository/in_memory_bd/entity"
)

func pointToTime(t time.Time) *time.Time {
	return &t
}

func TestCache_Get(t *testing.T) {
	type fields struct {
		data map[string]inMemoryBd.UrlInBd
		mtx  sync.RWMutex
	}
	type args struct {
		in0 context.Context
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.UrlInfo
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			fields: fields{
				data: map[string]inMemoryBd.UrlInBd{"google": inMemoryBd.UrlInBd{
					URL:       "google",
					Duration:  100,
					Headers:   []string{"x-1", "x-2"},
					LastCheck: 121,
					Status:    1,
				}},
				mtx: sync.RWMutex{},
			},
			args: args{
				in0: context.Background(),
				url: "google",
			},
			want: entity.UrlInfo{
				URL:       "google",
				Duration:  time.Millisecond * 100,
				Headers:   []string{"x-1", "x-2"},
				LastCheck: pointToTime(time.Unix(121, 0)),
				Status:    entity.Available,
			},
			wantErr: assert.NoError,
		},
		{
			name: "has error",
			fields: fields{
				data: map[string]inMemoryBd.UrlInBd{"google": inMemoryBd.UrlInBd{
					URL:       "google",
					Duration:  100,
					Headers:   []string{"x-1", "x-2"},
					LastCheck: 121,
					Status:    1,
				}},
				mtx: sync.RWMutex{},
			},
			args: args{
				in0: context.Background(),
				url: "google1",
			},
			want: entity.UrlInfo{},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, entity.NoDataErr, err.Error())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				data: tt.fields.data,
				mtx:  tt.fields.mtx,
			}
			got, err := c.Get(tt.args.in0, tt.args.url)
			if !tt.wantErr(t, err, fmt.Sprintf("GetStatus(%v, %v)", tt.args.in0, tt.args.url)) {
				return
			}

			if !assert.Equal(t, tt.want, got) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
