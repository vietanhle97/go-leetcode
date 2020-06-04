package Greatest_Sum_Divisible_By_Three_1262

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxDivisibleByThree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "maxDivisibleByThree",
			args: args{
				nums: []int{1, 2, 3, 4, 4},
			},
			want: 12,
		},
		{
			name: "maxDivisibleByThree",
			args: args{
				nums: []int{4},
			},
			want: 0,
		},
		{
			name: "maxDivisibleByThree",
			args: args{
				nums: []int{3, 6, 5, 1, 8},
			},
			want: 18,
		},
		{
			name: "maxDivisibleByThree",
			args: args{
				nums: []int{366, 809, 6, 792, 822, 181, 210, 588, 344, 618, 341, 410, 121, 864, 191, 749, 637, 169, 123, 472, 358, 908, 235, 914, 322, 946, 738, 754, 908, 272, 267, 326, 587, 267, 803, 281, 586, 707, 94, 627, 724, 469, 568, 57, 103, 984, 787, 552, 14, 545, 866, 494, 263, 157, 479, 823, 835, 100, 495, 773, 729, 921, 348, 871, 91, 386, 183, 979, 716, 806, 639, 290, 612, 322, 289, 910, 484, 300, 195, 546, 499, 213, 8, 623, 490, 473, 603, 721, 793, 418, 551, 331, 598, 670, 960, 483, 154, 317, 834, 352},
			},
			want: 50487,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxSumDivThree(tt.args.nums))
		})
	}
}
