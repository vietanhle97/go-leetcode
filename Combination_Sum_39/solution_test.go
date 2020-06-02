package Combination_Sum_39

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "combination 2",
			args: args{
				candidates: []int{3, 33, 333},
				target:     36,
			},
			want: [][]int{{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}, {3, 33}},
		},
		{
			name: "combination 2",
			args: args{
				candidates: []int{10, 1, 2, 7, 6, 1, 5},
				target:     8,
			},
			want: [][]int{{1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 2}, {1, 1, 1, 1, 2, 2}, {1, 1, 1, 5}, {1, 1, 2, 2, 2}, {1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 2, 2, 2}, {2, 6}},
		},
		{
			name: "combination 2",
			args: args{
				candidates: []int{1},
				target:     0,
			},
			want: [][]int{[]int{}},
		},
		{
			name: "combination 2",
			args: args{
				candidates: []int{2, 5, 2, 1, 2},
				target:     5,
			},
			want: [][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 2}, {1, 2, 2}, {5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, combinationSum(tt.args.candidates, tt.args.target))
		})
	}
}
