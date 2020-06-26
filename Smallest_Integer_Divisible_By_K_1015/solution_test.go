package Smallest_Integer_Divisible_By_K_1015

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_smallestIntegerDivisibleByK(t *testing.T) {
	type args struct {
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smallest integer divisible by K",
			args: args{
				K: 19,
			},
			want: 18,
		},
		{
			name: "smallest integer divisible by K",
			args: args{
				K: 23,
			},
			want: 22,
		},
		{
			name: "smallest integer divisible by K",
			args: args{
				K: 203,
			},
			want: 84,
		},
		{
			name: "smallest integer divisible by K",
			args: args{
				K: 10003,
			},
			want: 1428,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, smallestRepunitDivByK(tt.args.K))
		})
	}
}
