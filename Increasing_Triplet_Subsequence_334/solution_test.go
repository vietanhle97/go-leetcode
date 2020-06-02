package Increasing_Triplet_Subsequence_334

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_tripletIncreasingSub(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "triplet increaing subsequence",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: true,
		},
		{
			name: "triplet increaing subsequence",
			args: args{
				nums: []int{5, 4, 3, 2, 1},
			},
			want: false,
		},
		{
			name: "triplet increaing subsequence",
			args: args{
				nums: []int{1, 5, 3, 6, 2, 6, 8, 3, 2, 5, 2, 5, 2, 5, 6, 7, 8, 5, 34, 2, 2, 5, 7, 8},
			},
			want: true,
		},
		{
			name: "triplet increaing subsequence",
			args: args{
				nums: []int{5, 4, 3, 1, 7, 6},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, increasingTriplet(tt.args.nums))
		})
	}
}
