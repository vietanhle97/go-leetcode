package Number_of_Dice_Rolls_With_Target_Sum_1155

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_diceRollToTarget(t *testing.T) {
	type args struct {
		dices  int
		faces  int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "dices roll to target",
			args: args{
				dices:  1,
				faces:  6,
				target: 3,
			},
			want: 1,
		},
		{
			name: "dices roll to target",
			args: args{
				dices:  2,
				faces:  6,
				target: 7,
			},
			want: 6,
		},
		{
			name: "dices roll to target",
			args: args{
				dices:  6,
				faces:  6,
				target: 30,
			},
			want: 456,
		},
		{
			name: "dices roll to target",
			args: args{
				dices:  20,
				faces:  20,
				target: 60,
			},
			want: 26875834,
		},
		{
			name: "dices roll to target",
			args: args{
				dices:  4,
				faces:  8,
				target: 20,
			},
			want: 315,
		},
		{
			name: "dices roll to target",
			args: args{
				dices:  15,
				faces:  15,
				target: 36,
			},
			want: 319377986,
		},
		{
			name: "dices roll to target",
			args: args{
				dices:  30,
				faces:  30,
				target: 90,
			},
			want: 405749228,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, numRollsToTarget(tt.args.dices, tt.args.faces, tt.args.target))
		})
	}
}
