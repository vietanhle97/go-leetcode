package Shortest_Path_With_Alternating_Colors_1129

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_shortestPathDifferentColor(t *testing.T) {
	type args struct {
		n          int
		red_edges  [][]int
		blue_edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "shortest path",
			args: args{
				n:          5,
				red_edges:  [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
				blue_edges: [][]int{{1, 2}, {2, 3}, {3, 1}},
			},
			want: []int{0, 1, 2, 3, 7},
		},
		{
			name: "shortest path",
			args: args{
				n:          3,
				red_edges:  [][]int{{0, 1}, {1, 2}},
				blue_edges: [][]int{},
			},
			want: []int{0, 1, -1},
		},
		{
			name: "shortest path",
			args: args{
				n:          3,
				red_edges:  [][]int{{1, 0}},
				blue_edges: [][]int{{2, 1}},
			},
			want: []int{0, -1, -1},
		},
		{
			name: "shortest path",
			args: args{
				n:          3,
				red_edges:  [][]int{{0, 1}},
				blue_edges: [][]int{{1, 2}},
			},
			want: []int{0, 1, 2},
		},
		{
			name: "shortest path",
			args: args{
				n:          3,
				red_edges:  [][]int{{0, 1}, {0, 2}},
				blue_edges: [][]int{{1, 0}},
			},
			want: []int{0, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, shortestAlternatingPaths(tt.args.n, tt.args.red_edges, tt.args.blue_edges))
		})
	}
}
