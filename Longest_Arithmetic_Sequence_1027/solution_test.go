package Longest_Arithmetic_Sequence_1027

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_coinChange(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "longest arithmetic subsequence",
			args: args{
				nums: []int{3, 6, 9, 12},
			},
			want: 4,
		},
		{
			name: "longest arithmetic subsequence",
			args: args{
				nums: []int{9, 4, 7, 2, 10},
			},
			want: 3,
		},
		{
			name: "longest arithmetic subsequence",
			args: args{
				nums: []int{20, 1, 15, 3, 10, 5, 8},
			},
			want: 4,
		},
		{
			name: "longest arithmetic subsequence",
			args: args{
				nums: []int{83, 20, 17, 43, 52, 78, 68, 45},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, longestArithSeqLength(tt.args.nums))
		})
	}
}
