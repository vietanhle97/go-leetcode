package Spiral_Matrix_III_885

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_spiralMatrix_III(t *testing.T) {
	type args struct {
		R  int
		C  int
		r0 int
		c0 int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "combination 2",
			args: args{
				R:  5,
				C:  6,
				r0: 1,
				c0: 4,
			},
			want: [][]int{{1, 4}, {1, 5}, {2, 5}, {2, 4}, {2, 3}, {1, 3}, {0, 3}, {0, 4}, {0, 5}, {3, 5}, {3, 4}, {3, 3}, {3, 2}, {2, 2}, {1, 2}, {0, 2}, {4, 5}, {4, 4}, {4, 3}, {4, 2}, {4, 1}, {3, 1}, {2, 1}, {1, 1}, {0, 1}, {4, 0}, {3, 0}, {2, 0}, {1, 0}, {0, 0}},
		},
		{
			name: "combination 2",
			args: args{
				R:  1,
				C:  4,
				r0: 0,
				c0: 0,
			},
			want: [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		},
		{
			name: "combination 2",
			args: args{
				R:  6,
				C:  7,
				r0: 3,
				c0: 2,
			},
			want: [][]int{{3, 2}, {3, 3}, {4, 3}, {4, 2}, {4, 1}, {3, 1}, {2, 1}, {2, 2}, {2, 3}, {2, 4}, {3, 4}, {4, 4}, {5, 4}, {5, 3}, {5, 2}, {5, 1}, {5, 0}, {4, 0}, {3, 0}, {2, 0}, {1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 5}, {3, 5}, {4, 5}, {5, 5}, {0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}, {1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, spiralMatrixIII(tt.args.R, tt.args.C, tt.args.r0, tt.args.c0))
		})
	}
}
