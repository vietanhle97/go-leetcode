package Friend_Circles_547

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_friendCircle(t *testing.T) {
	type args struct {
		M [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "friend circle",
			args: args{
				M: [][]int{{1, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {1, 1, 0, 0, 0, 1, 0, 1, 0, 0}, {0, 0, 1, 0, 1, 0, 0, 1, 0, 0}, {0, 0, 0, 1, 0, 0, 1, 0, 0, 0}, {0, 0, 1, 0, 1, 0, 1, 0, 0, 0}, {0, 1, 0, 0, 0, 1, 0, 1, 1, 0}, {0, 0, 0, 1, 1, 0, 1, 0, 0, 0}, {0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 1, 0, 0, 1, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
			},
			want: 2,
		},
		{
			name: "friend circle",
			args: args{
				M: [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}},
			},
			want: 1,
		},
		{
			name: "friend circle",
			args: args{
				M: [][]int{{1, 1, 0}, {1, 1, 1}, {0, 1, 1}},
			},
			want: 1,
		},
		{
			name: "friend circle",
			args: args{
				M: [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findCircleNum(tt.args.M))
		})
	}
}
