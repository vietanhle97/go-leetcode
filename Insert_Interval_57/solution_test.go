package Insert_Interval_57

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_insertInterval(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "insert interval",
			args: args{
				intervals:   [][]int{{1, 3}, {6, 9}},
				newInterval: []int{2, 5},
			},
			want: [][]int{{1, 5}, {6, 9}},
		},
		{
			name: "insert interval",
			args: args{
				intervals:   [][]int{},
				newInterval: []int{5, 7},
			},
			want: [][]int{{5, 7}},
		},
		{
			name: "insert interval",
			args: args{
				intervals:   [][]int{{1, 5}},
				newInterval: []int{2, 3},
			},
			want: [][]int{{1, 5}},
		},
		{
			name: "insert interval",
			args: args{
				intervals:   [][]int{{1, 5}},
				newInterval: []int{1, 7},
			},
			want: [][]int{{1, 7}},
		},
		{
			name: "insert interval",
			args: args{
				intervals:   [][]int{{1, 5}},
				newInterval: []int{0, 3},
			},
			want: [][]int{{0, 5}},
		},
		{
			name: "insert interval",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{4, 8},
			},
			want: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, insert(tt.args.intervals, tt.args.newInterval))
		})
	}
}
