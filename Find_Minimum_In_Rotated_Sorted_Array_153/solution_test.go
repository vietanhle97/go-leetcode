package Find_Minimum_In_Rotated_Sorted_Array_153

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_findMinRotatedArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find min in rotated array",
			args: args{
				nums: []int{4, 5, 6, 7, 0, 1, 2},
			},
			want: 0,
		},
		{
			name: "find min in rotated array",
			args: args{
				nums: []int{2, 3, 4, 5, 6, 7, 8, 0, 1},
			},
			want: 0,
		},
		{
			name: "find min in rotated array",
			args: args{
				nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 9},
			},
			want: 0,
		},
		{
			name: "find min in rotated array",
			args: args{
				nums: []int{4, 5, 6, 0, 1, 2, 7, 8, 9},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findMin(tt.args.nums))
		})
	}
}
