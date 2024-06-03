package mapping

import (
	"reflect"
	"testing"
	"time"

	entity "url-checker/internal/domain"
	inMemoryBd "url-checker/internal/repository/in_memory_bd/entity"
)

func pointToTime(t time.Time) *time.Time {
	return &t
}

func TestURLBdToURLInfoMapping(t *testing.T) {
	type args struct {
		in inMemoryBd.UrlInBd
	}
	tests := []struct {
		name string
		args args
		want entity.UrlInfo
	}{
		{
			name: "mapping",
			args: args{
				in: inMemoryBd.UrlInBd{
					URL:       "URL",
					Duration:  10,
					Headers:   []string{"X-url", "X-test"},
					LastCheck: 30,
					Status:    2,
				},
			},
			want: entity.UrlInfo{
				URL:       "URL",
				Duration:  time.Millisecond * 10,
				Headers:   []string{"X-url", "X-test"},
				LastCheck: pointToTime(time.Unix(30, 0)),
				Status:    entity.NotAvailable,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := URLBdToURLInfoMapping(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URLBdToURLInfoMapping() = %v, want %v", got, tt.want)
			}
		})
	}
}
