package Median_Of_Two_Sorted_Arrays_4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_medianTwoSortedArray(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "median of 2 sorted arrays",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: float64(2),
		},
		{
			name: "median of 2 sorted arrays",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.5,
		},
		{
			name: "median of 2 sorted arrays",
			args: args{
				nums1: []int{1, 2, 4, 5, 6, 7},
				nums2: []int{3, 4, 5, 6},
			},
			want: 4.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findMedianSortedArrays(tt.args.nums1, tt.args.nums2))
		})
	}
}
