package Super_Ugly_Number_313

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_superUglyNumber(t *testing.T) {
	type args struct {
		n      int
		primes []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "super ugly number",
			args: args{
				n:      1000,
				primes: []int{2, 7, 13, 19},
			},
			want: 85489664,
		},
		{
			name: "super ugly number",
			args: args{
				n:      12,
				primes: []int{2, 7, 13, 19},
			},
			want: 32,
		},
		{
			name: "super ugly number",
			args: args{
				n:      1000,
				primes: []int{2, 7, 13, 19, 21, 37, 31},
			},
			want: 342657,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, nthSuperUglyNumber(tt.args.n, tt.args.primes))
		})
	}
}
