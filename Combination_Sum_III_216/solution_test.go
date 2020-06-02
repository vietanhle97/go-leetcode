package Combination_Sum_III_216

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	type args struct {
		target int
		size   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "combination 3",
			args: args{
				target: 7,
				size:   3,
			},
			want: [][]int{{1, 2, 4}},
		},
		{
			name: "combination 3",
			args: args{
				target: 9,
				size:   3,
			},
			want: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		},
		{
			name: "combination 3",
			args: args{
				target: 16,
				size:   4,
			},
			want: [][]int{{1, 2, 4, 9}, {1, 2, 5, 8}, {1, 2, 6, 7}, {1, 3, 4, 8}, {1, 3, 5, 7}, {1, 4, 5, 6}, {2, 3, 4, 7}, {2, 3, 5, 6}},
		},
		{
			name: "combination 3",
			args: args{
				target: 15,
				size:   3,
			},
			want: [][]int{{1, 5, 9}, {1, 6, 8}, {2, 4, 9}, {2, 5, 8}, {2, 6, 7}, {3, 4, 8}, {3, 5, 7}, {4, 5, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, combinationSum3(tt.args.size, tt.args.target))
		})
	}
}
