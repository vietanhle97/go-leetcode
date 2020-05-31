package Coin_Change_322

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "coin change",
			args: args{
				coins:  []int{2, 5, 10, 1},
				amount: 27,
			},
			want: 4,
		},
		{
			name: "coin change",
			args: args{
				coins:  []int{2},
				amount: 1,
			},
			want: -1,
		},
		{
			name: "coin change",
			args: args{
				coins:  []int{1},
				amount: 0,
			},
			want: 0,
		},
		{
			name: "coin change",
			args: args{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			want: 3,
		},
		{
			name: "coin change",
			args: args{
				coins:  []int{1, 2147483647},
				amount: 2,
			},
			want: 2,
		},
		{
			name: "coin change",
			args: args{
				coins:  []int{2147483647},
				amount: 2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, coinChange(tt.args.coins, tt.args.amount))
		})
	}
}
