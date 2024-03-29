package Find_Peak_Element_162

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find peak element",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: 2,
		},
		{
			name: "find peak element",
			args: args{
				nums: []int{1, 2, 1, 3, 5, 6, 4},
			},
			want: 5,
		},
		{
			name: "find peak element",
			args: args{
				nums: []int{2, 1},
			},
			want: 0,
		},
		{
			name: "find peak element",
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
		{
			name: "find peak element",
			args: args{
				nums: []int{1, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findPeakElement(tt.args.nums))
		})
	}
}
