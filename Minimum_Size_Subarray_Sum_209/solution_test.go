package Minimum_Size_Subarray_Sum_209

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minSizeSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		s    int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "min size of sub array sum",
			args: args{
				nums: []int{2, 3, 1, 1, 4, 3},
				s:    7,
			},
			want: 2,
		},
		{
			name: "min size of sub array sum",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				s:    11,
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minSubArrayLen(tt.args.s, tt.args.nums))
		})
	}
}
