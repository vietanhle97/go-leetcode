package Combination_Sum_II_40

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
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
			want: [][]int{{3, 33}},
		},
		{
			name: "combination 2",
			args: args{
				candidates: []int{10, 1, 2, 7, 6, 1, 5},
				target:     8,
			},
			want: [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
		{
			name: "combination 2",
			args: args{
				candidates: []int{1},
				target:     0,
			},
			want: [][]int{{}},
		},
		{
			name: "combination 2",
			args: args{
				candidates: []int{2, 5, 2, 1, 2},
				target:     5,
			},
			want: [][]int{{1, 2, 2}, {5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, combinationSum2(tt.args.candidates, tt.args.target))
		})
	}
}
