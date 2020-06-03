package Partition_Equal_Subset_Sum_416

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
		want bool
	}{
		{
			name: "sort color",
			args: args{
				nums: []int{2, 0, 2, 1, 1, 0},
			},
			want: true,
		},
		{
			name: "sort color",
			args: args{
				nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100},
			},
			want: false,
		},
		{
			name: "sort color",
			args: args{
				nums: []int{1, 5, 11, 5},
			},
			want: true,
		},
		{
			name: "sort color",
			args: args{
				nums: []int{1, 2, 3, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, canPartition(tt.args.nums))
		})
	}
}
