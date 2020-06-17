package Rotting_Oranges_994

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rottingOranges(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "rotting oranges",
			args: args{
				grid: [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			},
			want: 4,
		},
		{
			name: "rotting oranges",
			args: args{
				grid: [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			},
			want: -1,
		},
		{
			name: "rotting oranges",
			args: args{
				grid: [][]int{{2, 2}, {1, 1}, {0, 0}, {2, 0}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, orangesRotting(tt.args.grid))
		})
	}
}
