package Spiral_Matrix_II_59

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_spiralMatrixII(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "spiral matrix",
			args: args{
				n: 3,
			},
			want: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
		},
		{
			name: "spiral matrix",
			args: args{
				n: 4,
			},
			want: [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}},
		},
		{
			name: "spiral matrix",
			args: args{
				n: 2,
			},
			want: [][]int{{1, 2}, {4, 3}},
		},
		{
			name: "spiral matrix",
			args: args{
				n: 5,
			},
			want: [][]int{{1, 2, 3, 4, 5}, {16, 17, 18, 19, 6}, {15, 24, 25, 20, 7}, {14, 23, 22, 21, 8}, {13, 12, 11, 10, 9}},
		},
		{
			name: "spiral matrix",
			args: args{
				n: 6,
			},
			want: [][]int{{1, 2, 3, 4, 5, 6}, {20, 21, 22, 23, 24, 7}, {19, 32, 33, 34, 25, 8}, {18, 31, 36, 35, 26, 9}, {17, 30, 29, 28, 27, 10}, {16, 15, 14, 13, 12, 11}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, generateMatrix(tt.args.n))
		})
	}
}
