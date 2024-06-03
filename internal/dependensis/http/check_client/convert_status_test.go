package check_client

import (
	"testing"

	entity "url-checker/internal/domain"
)

func Test_convertStatus(t *testing.T) {
	type args struct {
		status int
	}
	tests := []struct {
		name string
		args args
		want entity.Status
	}{
		{
			name: "200",
			args: args{
				status: 200,
			},
			want: entity.Available,
		},
		{
			name: "230",
			args: args{
				status: 230,
			},
			want: entity.Available,
		},
		{
			name: "300",
			args: args{
				status: 300,
			},
			want: entity.Moved,
		},
		{
			name: "350",
			args: args{
				status: 350,
			},
			want: entity.Moved,
		},
		{
			name: "400",
			args: args{
				status: 400,
			},
			want: entity.NotAvailable,
		},
		{
			name: "404",
			args: args{
				status: 404,
			},
			want: entity.NotAvailable,
		},
		{
			name: "500",
			args: args{
				status: 500,
			},
			want: entity.NotAvailable,
		},
		{
			name: "510",
			args: args{
				status: 510,
			},
			want: entity.NotAvailable,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertStatus(tt.args.status); got != tt.want {
				t.Errorf("convertStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
