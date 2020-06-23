package Find_First_And_Last_Position_Of_Element_In_Sorted_Array_34

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findFirstAndLastPositionOfTarget(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "find first and last position of element in sorted array",
			args: args{
				nums:   []int{1, 2, 2, 3, 3, 4, 4, 5, 6, 6, 7, 7, 7, 7, 7, 7, 8, 8, 9, 9, 9, 10},
				target: 7,
			},
			want: []int{10, 15},
		},
		{
			name: "find first and last position of element in sorted array",
			args: args{
				nums:   []int{1, 2, 2, 3, 3, 4, 4, 5, 6, 6, 7, 7, 7, 7, 7, 7, 8, 8, 9, 9, 9, 10},
				target: 3,
			},
			want: []int{3, 4},
		},
		{
			name: "find first and last position of element in sorted array",
			args: args{
				nums:   []int{1, 2, 2, 3, 3, 4, 4, 5, 6, 6, 7, 7, 7, 7, 7, 7, 8, 8, 9, 9, 9, 10},
				target: 9,
			},
			want: []int{18, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, searchRange(tt.args.nums, tt.args.target))
		})
	}
}
