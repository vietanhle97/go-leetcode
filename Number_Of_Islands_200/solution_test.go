package Number_Of_Islands_200

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numOfIsland(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "number of islands",
			args: args{
				grid: [][]byte{{'1', '0', '1', '1', '1'}, {'1', '0', '1', '0', '1'}, {'1', '1', '1', '0', '1'}},
			},
			want: 1,
		},
		{
			name: "number of islands",
			args: args{
				grid: [][]byte{{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, numIslands(tt.args.grid))
		})
	}
}
