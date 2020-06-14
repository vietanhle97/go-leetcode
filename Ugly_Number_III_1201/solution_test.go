package Ugly_Number_III_1201

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_uglyNumberIII(t *testing.T) {
	type args struct {
		n int
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ugly number III",
			args: args{
				n: 1000,
				a: 5,
				b: 100,
				c: 113,
			},
			want: 4830,
		},
		{
			name: "ugly number III",
			args: args{
				n: 1000000000,
				a: 2,
				b: 217983653,
				c: 336916467,
			},
			want: 1999999984,
		},
		{
			name: "ugly number III",
			args: args{
				n: 1000000000,
				a: 2,
				b: 3,
				c: 37,
			},
			want: 1480000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, nthUglyNumber(tt.args.n, tt.args.a, tt.args.b, tt.args.c))
		})
	}
}
