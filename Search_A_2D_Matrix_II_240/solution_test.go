package Search_A_2D_Matrix_II_240

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchIn2DMatrixII(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "search in 2D matrix",
			args: args{
				matrix: [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
				target: 5,
			},
			want: true,
		},
		{
			name: "search in 2D matrix",
			args: args{
				matrix: [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
				target: 20,
			},
			want: false,
		},
		{
			name: "search in 2D matrix",
			args: args{
				matrix: [][]int{{-5}},
				target: -5,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, searchMatrix(tt.args.matrix, tt.args.target))
		})
	}
}
