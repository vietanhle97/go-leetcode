package Merge_Intervals_56

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_intervalsGame(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "intervals",
			args: args{
				intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: "intervals",
			args: args{
				intervals: [][]int{{1, 4}, {5, 6}},
			},
			want: [][]int{{1, 4}, {5, 6}},
		},
		{
			name: "intervals",
			args: args{
				intervals: [][]int{{1, 4}, {4, 5}},
			},
			want: [][]int{{1, 5}},
		},
		{
			name: "intervals",
			args: args{
				intervals: [][]int{{1, 4}, {2, 5}},
			},
			want: [][]int{{1, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, merge(tt.args.intervals))
		})
	}
}
