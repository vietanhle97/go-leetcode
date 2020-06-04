package _Sum_With_Multiplicity_923

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_3SumMulti(t *testing.T) {
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
			name: "3 sum mutiply",
			args: args{
				nums:   []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
				target: 8,
			},
			want: 20,
		},
		{
			name: "3 sum mutiply",
			args: args{
				nums:   []int{1, 1, 2, 2, 2, 2},
				target: 5,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, threeSumMulti(tt.args.nums, tt.args.target))
		})
	}
}
