package Subarray_Sum_Equals_K_560

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "combination 2",
			args: args{
				nums:   []int{-1, -1, 1},
				target: 1,
			},
			want: 1,
		},
		{
			name: "combination 2",
			args: args{
				nums:   []int{10, 1, 2, 7, 6, 1, 5},
				target: 8,
			},
			want: 0,
		},
		{
			name: "combination 2",
			args: args{
				nums:   []int{28, 54, 7, -70, 22, 65, -6},
				target: 100,
			},
			want: 1,
		},
		{
			name: "combination 2",
			args: args{
				nums:   []int{2, 5, 2, 1, 2},
				target: 5,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, subarraySum(tt.args.nums, tt.args.target))
		})
	}
}
