package Contains_Duplicate_III_220

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_containDuplicateIII(t *testing.T) {
	type args struct {
		nums []int
		k    int
		t    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "contains duplicate III",
			args: args{
				nums: []int{1, 2, 3, 1},
				k:    3,
				t:    0,
			},
			want: true,
		},
		{
			name: "contains duplicate III",
			args: args{
				nums: []int{1, 0, 1, 1},
				k:    1,
				t:    2,
			},
			want: true,
		},
		{
			name: "contains duplicate III",
			args: args{
				nums: []int{1, 5, 9, 1, 5, 9},
				k:    2,
				t:    3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, containsNearbyAlmostDuplicate(tt.args.nums, tt.args.k, tt.args.t))
		})
	}
}
