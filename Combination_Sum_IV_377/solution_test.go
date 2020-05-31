package Combination_Sum_IV_377

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_combinationSum4(t *testing.T) {
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
			name: "combination 4",
			args: args{
				nums:   []int{3, 33, 333},
				target: 10000,
			},
			want: 0,
		},
		{
			name: "combination 4",
			args: args{
				nums:   []int{1, 2, 3},
				target: 4,
			},
			want: 7,
		},
		{
			name: "combination 4",
			args: args{
				nums:   []int{1},
				target: 0,
			},
			want: 0,
		},
		{
			name: "combination 4",
			args: args{
				nums:   []int{1, 2, 5},
				target: 11,
			},
			want: 218,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, combinationSum4(tt.args.nums, tt.args.target))
		})
	}
}
