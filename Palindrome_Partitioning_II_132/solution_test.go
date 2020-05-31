package Palindrome_Partitioning_II_132

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_palindromePartition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "min Cut",
			args: args{
				s: "aab",
			},
			want: 1,
		},
		{
			name: "min Cut",
			args: args{
				s: "abbab",
			},
			want: 1,
		},
		{
			name: "min Cut",
			args: args{
				s: "fifgbeajcacehiicccfecbfhhgfiiecdcjjffbghdidbhbdbfbfjccgbbdcjheccfbhafehieabbdfeigbiaggchaeghaijfbjhi",
			},
			want: 75,
		},
		{
			name: "min Cut",
			args: args{
				s: "efe",
			},
			want: 0,
		},
		{
			name: "min Cut",
			args: args{
				s: "abcd",
			},
			want: 3,
		},
		{
			name: "min Cut",
			args: args{
				s: "aabbc",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minCut(tt.args.s))
		})
	}
}
