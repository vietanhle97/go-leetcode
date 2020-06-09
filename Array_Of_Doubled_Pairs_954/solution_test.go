package Array_Of_Doubled_Pairs_954

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_arrayDoubledPairs(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "array of doubled pairs",
			args: args{
				nums: []int{7, -15, -15, 23, -3, 80, -35, 40, 68, 22, 44, 98, 20, 0, -34, 8, 40, 41, 16, 46, 16, 49, -6, -11, 35, -15, -74, 72, -8, 60, 40, -2, 0, -6, 34, 14, -16, -92, 54, 14, -68, 82, -30, 50, 22, 25, 16, 70, -1, -96, 11, 45, 54, 40, 92, -35, 29, 80, 46, -30, 27, 7, -70, -37, 41, -46, -98, 1, -33, -24, -86, -70, 80, -43, 98, -49, 30, 0, 27, 2, 82, 36, 0, -48, 3, -100, 58, 32, 90, -22, -50, -12, 36, 6, -3, -66, 72, 8, 49, -30},
			},
			want: true,
		},
		{
			name: "array of doubled pairs",
			args: args{
				nums: []int{2, 1, 2, 1, 1, 1, 2, 2},
			},
			want: true,
		},
		{
			name: "array of doubled pairs",
			args: args{
				nums: []int{-62, 86, 96, -18, 43, -24, -76, 13, -31, -26, -88, -13, 96, -44, 9, -20, -42, 100, -44, -24, -36, 28, -32, 58, -72, 20, 48, -36, -45, 14, 24, -64, -90, -44, -16, 86, -33, 48, 26, 29, -84, 10, 46, 50, -66, -8, -38, 92, -19, 43, 48, -38, -22, 18, -32, -48, -64, -10, -22, -48},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, canReorderDoubled(tt.args.nums))
		})
	}
}
