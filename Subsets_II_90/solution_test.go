package Subsets_II_90

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_subSetII(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "subset",
			args: args{
				nums: []int{1, 2, 1, 2, 4, 3, 2, 1, 2, 3},
			},
			want: [][]int{{}, {1}, {1, 1}, {1, 1, 1}, {1, 1, 1, 2}, {1, 1, 1, 2, 2}, {1, 1, 1, 2, 2, 2}, {1, 1, 1, 2, 2, 2, 2}, {1, 1, 1, 2, 2, 2, 2, 3}, {1, 1, 1, 2, 2, 2, 2, 3, 3}, {1, 1, 1, 2, 2, 2, 2, 3, 3, 4}, {1, 1, 1, 2, 2, 2, 2, 3, 4}, {1, 1, 1, 2, 2, 2, 2, 4}, {1, 1, 1, 2, 2, 2, 3}, {1, 1, 1, 2, 2, 2, 3, 3}, {1, 1, 1, 2, 2, 2, 3, 3, 4}, {1, 1, 1, 2, 2, 2, 3, 4}, {1, 1, 1, 2, 2, 2, 4}, {1, 1, 1, 2, 2, 3}, {1, 1, 1, 2, 2, 3, 3}, {1, 1, 1, 2, 2, 3, 3, 4}, {1, 1, 1, 2, 2, 3, 4}, {1, 1, 1, 2, 2, 4}, {1, 1, 1, 2, 3}, {1, 1, 1, 2, 3, 3}, {1, 1, 1, 2, 3, 3, 4}, {1, 1, 1, 2, 3, 4}, {1, 1, 1, 2, 4}, {1, 1, 1, 3}, {1, 1, 1, 3, 3}, {1, 1, 1, 3, 3, 4}, {1, 1, 1, 3, 4}, {1, 1, 1, 4}, {1, 1, 2}, {1, 1, 2, 2}, {1, 1, 2, 2, 2}, {1, 1, 2, 2, 2, 2}, {1, 1, 2, 2, 2, 2, 3}, {1, 1, 2, 2, 2, 2, 3, 3}, {1, 1, 2, 2, 2, 2, 3, 3, 4}, {1, 1, 2, 2, 2, 2, 3, 4}, {1, 1, 2, 2, 2, 2, 4}, {1, 1, 2, 2, 2, 3}, {1, 1, 2, 2, 2, 3, 3}, {1, 1, 2, 2, 2, 3, 3, 4}, {1, 1, 2, 2, 2, 3, 4}, {1, 1, 2, 2, 2, 4}, {1, 1, 2, 2, 3}, {1, 1, 2, 2, 3, 3}, {1, 1, 2, 2, 3, 3, 4}, {1, 1, 2, 2, 3, 4}, {1, 1, 2, 2, 4}, {1, 1, 2, 3}, {1, 1, 2, 3, 3}, {1, 1, 2, 3, 3, 4}, {1, 1, 2, 3, 4}, {1, 1, 2, 4}, {1, 1, 3}, {1, 1, 3, 3}, {1, 1, 3, 3, 4}, {1, 1, 3, 4}, {1, 1, 4}, {1, 2}, {1, 2, 2}, {1, 2, 2, 2}, {1, 2, 2, 2, 2}, {1, 2, 2, 2, 2, 3}, {1, 2, 2, 2, 2, 3, 3}, {1, 2, 2, 2, 2, 3, 3, 4}, {1, 2, 2, 2, 2, 3, 4}, {1, 2, 2, 2, 2, 4}, {1, 2, 2, 2, 3}, {1, 2, 2, 2, 3, 3}, {1, 2, 2, 2, 3, 3, 4}, {1, 2, 2, 2, 3, 4}, {1, 2, 2, 2, 4}, {1, 2, 2, 3}, {1, 2, 2, 3, 3}, {1, 2, 2, 3, 3, 4}, {1, 2, 2, 3, 4}, {1, 2, 2, 4}, {1, 2, 3}, {1, 2, 3, 3}, {1, 2, 3, 3, 4}, {1, 2, 3, 4}, {1, 2, 4}, {1, 3}, {1, 3, 3}, {1, 3, 3, 4}, {1, 3, 4}, {1, 4}, {2}, {2, 2}, {2, 2, 2}, {2, 2, 2, 2}, {2, 2, 2, 2, 3}, {2, 2, 2, 2, 3, 3}, {2, 2, 2, 2, 3, 3, 4}, {2, 2, 2, 2, 3, 4}, {2, 2, 2, 2, 4}, {2, 2, 2, 3}, {2, 2, 2, 3, 3}, {2, 2, 2, 3, 3, 4}, {2, 2, 2, 3, 4}, {2, 2, 2, 4}, {2, 2, 3}, {2, 2, 3, 3}, {2, 2, 3, 3, 4}, {2, 2, 3, 4}, {2, 2, 4}, {2, 3}, {2, 3, 3}, {2, 3, 3, 4}, {2, 3, 4}, {2, 4}, {3}, {3, 3}, {3, 3, 4}, {3, 4}, {4}},
		},
		{
			name: "subset",
			args: args{
				nums: []int{1, 2, 1, 2, 4, 3},
			},
			want: [][]int{{}, {1}, {1, 1}, {1, 1, 2}, {1, 1, 2, 2}, {1, 1, 2, 2, 3}, {1, 1, 2, 2, 3, 4}, {1, 1, 2, 2, 4}, {1, 1, 2, 3}, {1, 1, 2, 3, 4}, {1, 1, 2, 4}, {1, 1, 3}, {1, 1, 3, 4}, {1, 1, 4}, {1, 2}, {1, 2, 2}, {1, 2, 2, 3}, {1, 2, 2, 3, 4}, {1, 2, 2, 4}, {1, 2, 3}, {1, 2, 3, 4}, {1, 2, 4}, {1, 3}, {1, 3, 4}, {1, 4}, {2}, {2, 2}, {2, 2, 3}, {2, 2, 3, 4}, {2, 2, 4}, {2, 3}, {2, 3, 4}, {2, 4}, {3}, {3, 4}, {4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, subsetsWithDup(tt.args.nums))
		})
	}
}
