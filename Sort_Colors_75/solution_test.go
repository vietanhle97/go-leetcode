package Sort_Colors_75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_asteroidCollision(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sort color",
			args: args{
				nums: []int{2, 0, 2, 1, 1, 0},
			},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "sort color",
			args: args{
				nums: []int{2, 0, 2, 1, 1, 0, 1, 2, 0, 1, 2, 0, 2, 2, 1, 2, 0},
			},
			want: []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2},
		},
		{
			name: "sort color",
			args: args{
				nums: []int{2, 1, 0},
			},
			want: []int{0, 1, 2},
		},
		{
			name: "sort color",
			args: args{
				nums: []int{2, 1, 2, 2, 0, 1},
			},
			want: []int{0, 1, 1, 2, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			assert.Equal(t, tt.want, tt.args.nums)
		})
	}
}
