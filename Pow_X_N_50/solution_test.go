package Pow_x__n__50

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "find x power to the n",
			args: args{
				x: 2.0,
				n: 10,
			},
			want: 1024.0,
		},
		{
			name: "find x power to the n",
			args: args{
				x: 0.00001,
				n: 2147483647,
			},
			want: 0.0,
		},
		{
			name: "find x power to the n",
			args: args{
				x: 2.1,
				n: 3,
			},
			want: 9.261000000000001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, myPow(tt.args.x, tt.args.n))
		})
	}
}
