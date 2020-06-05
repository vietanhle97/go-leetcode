package Binary_Subarrays_With_Sum_930

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_binarySubWithSum(t *testing.T) {
	type args struct {
		A []int
		S int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "binary subarray with sum",
			args: args{
				A: []int{1, 0, 1, 0, 1},
				S: 2,
			},
			want: 4,
		},
		{
			name: "binary subarray with sum",
			args: args{
				A: []int{1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1},
				S: 5,
			},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, numSubarraysWithSum(tt.args.A, tt.args.S))
		})
	}
}
