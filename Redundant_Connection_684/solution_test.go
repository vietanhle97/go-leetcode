package Redundant_Connection_684

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_redundantConn(t *testing.T) {
	type args struct {
		edges [][]int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "redundant connection",
			args: args{
				edges: [][]int{{1, 2}, {1, 3}, {2, 3}},
			},
			want: []int{2, 3},
		},
		{
			name: "redundant connection",
			args: args{
				edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
			},
			want: []int{1, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findRedundantConnection(tt.args.edges))
		})
	}
}
