package Search_A_2D_Matrix_74

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchIn2DMatrix(t *testing.T) {
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
				matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}},
				target: 13,
			},
			want: false,
		},
		{
			name: "search in 2D matrix",
			args: args{
				matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}},
				target: 3,
			},
			want: true,
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
