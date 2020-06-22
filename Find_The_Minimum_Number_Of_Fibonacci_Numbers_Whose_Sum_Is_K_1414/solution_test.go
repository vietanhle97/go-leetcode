package Find_The_Minimum_Number_Of_Fibonacci_Numbers_Whose_Sum_Is_K_1414

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fibSumToK(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "minimum fib number sum to k",
			args: args{
				k: 1000,
			},
			want: 2,
		},
		{
			name: "minimum fib number sum to k",
			args: args{
				k: 1000000000,
			},
			want: 14,
		},
		{
			name: "minimum fib number sum to k",
			args: args{
				k: 9111,
			},
			want: 7,
		},
		{
			name: "minimum fib number sum to k",
			args: args{
				k: 9999,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findMinFibonacciNumbers(tt.args.k))
		})
	}
}
