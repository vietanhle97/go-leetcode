package Sum_Of_Distances_In_Tree_834

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sumDistanceTree(t *testing.T) {
	type args struct {
		N     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sum of distances in tree",
			args: args{
				N:     6,
				edges: [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}},
			},
			want: []int{8, 12, 6, 10, 10, 10},
		},
		{
			name: "sum of distances in tree",
			args: args{
				N:     6,
				edges: [][]int{{0, 5}, {0, 2}, {2, 3}, {2, 4}, {2, 1}},
			},
			want: []int{8, 10, 6, 10, 10, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, sumOfDistancesInTree(tt.args.N, tt.args.edges))
		})
	}
}
