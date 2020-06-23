package Find_The_Smallest_Divisor_Given_A_Threshold_1283

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_findSmallestDivisorWithThreshold(t *testing.T) {
	type args struct {
		nums      []int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smallest divisor given a threshold",
			args: args{
				nums:      []int{1, 2, 5, 9},
				threshold: 6,
			},
			want: 5,
		},
		{
			name: "smallest divisor given a threshold",
			args: args{
				nums:      []int{19},
				threshold: 5,
			},
			want: 5,
		},
		{
			name: "smallest divisor given a threshold",
			args: args{
				nums:      []int{2, 3, 5, 7, 11},
				threshold: 11,
			},
			want: 3,
		},
		{
			name: "smallest divisor given a threshold",
			args: args{
				nums:      []int{1, 2, 5, 9},
				threshold: 4,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, smallestDivisor(tt.args.nums, tt.args.threshold))
		})
	}
}
