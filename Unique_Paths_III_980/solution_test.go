package Unique_Paths_III_980

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_uniquepathIII(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "unique path",
			args: args{
				grid: [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}},
			},
			want: 2,
		},
		{
			name: "unique path",
			args: args{
				grid: [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 2}},
			},
			want: 4,
		},
		{
			name: "unique path",
			args: args{
				grid: [][]int{{0, 1}, {2, 0}},
			},
			want: 0,
		},
		{
			name: "unique path",
			args: args{
				grid: [][]int{{1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 2}},
			},
			want: 16,
		},
		{
			name: "unique path",
			args: args{
				grid: [][]int{{1, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 2}},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, uniquePathsIII(tt.args.grid))
		})
	}
}
