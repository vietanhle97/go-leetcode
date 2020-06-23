package Capacity_To_Ship_Packages_Within_D_Days_1011

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_capacityToShipPackages(t *testing.T) {
	type args struct {
		weights []int
		D       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "minimum capacity to ship packages within D days",
			args: args{
				weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				D:       5,
			},
			want: 15,
		},
		{
			name: "minimum capacity to ship packages within D days",
			args: args{
				weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				D:       1,
			},
			want: 55,
		},
		{
			name: "minimum capacity to ship packages within D days",
			args: args{
				weights: []int{3, 2, 2, 4, 1, 4},
				D:       3,
			},
			want: 6,
		},
		{
			name: "minimum capacity to ship packages within D days",
			args: args{
				weights: []int{1, 2, 3, 1, 1},
				D:       4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, shipWithinDays(tt.args.weights, tt.args.D))
		})
	}
}
