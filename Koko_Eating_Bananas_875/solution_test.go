package Koko_Eating_Bananas_875

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_kokoEatingBanana(t *testing.T) {
	type args struct {
		piles []int
		H     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "minimum hours to eat all bananas",
			args: args{
				piles: []int{3, 6, 7, 11},
				H:     8,
			},
			want: 4,
		},
		{
			name: "minimum hours to eat all bananas",
			args: args{
				piles: []int{30, 11, 23, 4, 20},
				H:     5,
			},
			want: 30,
		},
		{
			name: "minimum hours to eat all bananas",
			args: args{
				piles: []int{30, 11, 23, 4, 20},
				H:     6,
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minEatingSpeed(tt.args.piles, tt.args.H))
		})
	}
}
