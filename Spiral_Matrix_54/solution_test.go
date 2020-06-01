package Spiral_Matrix_54

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_spiralMatrix(t *testing.T) {
	type args struct {
		spiral [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "spiral matrix",
			args: args{
				spiral: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "spiral matrix",
			args: args{
				spiral: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name: "spiral matrix",
			args: args{
				spiral: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "spiral matrix",
			args: args{
				spiral: [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			name: "spiral matrix",
			args: args{
				spiral: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			},
			want: []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, spiralOrder(tt.args.spiral))
		})
	}
}
