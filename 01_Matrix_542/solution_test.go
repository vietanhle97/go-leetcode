package _1_Matrix_542

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_matrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "matrix",
			args: args{
				matrix: [][]int{{1, 1, 1}, {0, 1, 0}, {1, 1, 1}},
			},
			want: [][]int{{1, 2, 1}, {0, 1, 0}, {1, 2, 1}},
		},
		{
			name: "matrix",
			args: args{
				matrix: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}},
			},
			want: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, updateMatrix(tt.args.matrix))
		})
	}
}
