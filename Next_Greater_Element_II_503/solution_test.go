package Next_Greater_Element_II_503

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NextGreaterII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "next greater II",
			args: args{
				nums: []int{1, 2, 1},
			},
			want: []int{2, -1, 2},
		},
		{
			name: "next greater II",
			args: args{
				nums: []int{122, 209, 1698, 1822, 17, 3071, 9634, 9741, 9794, 487, 24551, 24641, 24223, 28300, 28410, 28510, 28552},
			},
			want: []int{209, 1698, 1822, 3071, 3071, 9634, 9741, 9794, 24551, 24551, 24641, 28300, 28300, 28410, 28510, 28552, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, nextGreaterElements(tt.args.nums))
		})
	}
}
