package _32_Pattern_456

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_find132pattern(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "132 pattern",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: false,
		},
		{
			name: "132 pattern",
			args: args{
				nums: []int{3, 1, 4, 2},
			},
			want: true,
		},
		{
			name: "132 pattern",
			args: args{
				nums: []int{-1, 3, 2, 0},
			},
			want: true,
		},
		{
			name: "132 pattern",
			args: args{
				nums: []int{-2, 1, -1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, find132pattern(tt.args.nums))
		})
	}
}
