package Find_K_Closest_Elements_658

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findKClosestElements(t *testing.T) {
	type args struct {
		arr []int
		k   int
		x   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "find k closest elements",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7},
				k:   4,
				x:   3,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "find k closest elements",
			args: args{
				arr: []int{0, 1, 1, 1, 2, 3, 6, 7, 8, 9},
				k:   9,
				x:   4,
			},
			want: []int{0, 1, 1, 1, 2, 3, 6, 7, 8},
		},
		{
			name: "find k closest elements",
			args: args{
				arr: []int{1, 1, 1, 10, 10, 10},
				k:   1,
				x:   9,
			},
			want: []int{10},
		},
		{
			name: "find k closest elements",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   -1,
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findClosestElements(tt.args.arr, tt.args.k, tt.args.x))

		})
	}
}
