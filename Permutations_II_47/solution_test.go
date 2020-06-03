package Permutations_II_47

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_permutations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "permutations",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
		{
			name: "permutations",
			args: args{
				nums: []int{1, 1, 2, 3},
			},
			want: [][]int{{1, 1, 2, 3}, {1, 1, 3, 2}, {1, 2, 1, 3}, {1, 2, 3, 1}, {1, 3, 1, 2}, {1, 3, 2, 1}, {2, 1, 1, 3}, {2, 1, 3, 1}, {2, 3, 1, 1}, {3, 1, 1, 2}, {3, 1, 2, 1}, {3, 2, 1, 1}},
		},
		{
			name: "permutations",
			args: args{
				nums: []int{2, 2, 3, 1},
			},
			want: [][]int{{2, 2, 3, 1}, {2, 2, 1, 3}, {2, 3, 2, 1}, {2, 3, 1, 2}, {2, 1, 2, 3}, {2, 1, 3, 2}, {3, 2, 2, 1}, {3, 2, 1, 2}, {3, 1, 2, 2}, {1, 2, 2, 3}, {1, 2, 3, 2}, {1, 3, 2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, permuteUnique(tt.args.nums))
		})
	}
}
