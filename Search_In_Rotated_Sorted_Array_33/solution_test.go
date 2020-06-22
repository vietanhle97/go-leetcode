package Search_In_Rotated_Sorted_Array_33

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_searchRotatedArray(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search in rotated array",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
		{
			name: "search in rotated array",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: -1,
		},
		{
			name: "search in rotated array",
			args: args{
				nums:   []int{0, 1, 2, 3, 4, 5, 6, 7, 9},
				target: 7,
			},
			want: 7,
		},
		{
			name: "search in rotated array",
			args: args{
				nums:   []int{4, 5, 6, 0, 1, 2, 7, 8, 99},
				target: 2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, search(tt.args.nums, tt.args.target))
		})
	}
}
