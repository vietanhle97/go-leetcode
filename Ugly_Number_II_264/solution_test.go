package Ugly_Number_II_264

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_uglyNumberII(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ugly number II",
			args: args{
				n: 10,
			},
			want: 12,
		},
		{
			name: "ugly number II",
			args: args{
				n: 1690,
			},
			want: 2123366400,
		},
		{
			name: "ugly number II",
			args: args{
				n: 1000,
			},
			want: 51200000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, nthUglyNumber(tt.args.n))
		})
	}
}
