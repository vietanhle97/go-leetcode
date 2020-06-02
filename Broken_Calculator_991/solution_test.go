package Broken_Calculator_991

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_calculator(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "calculator",
			args: args{
				x: 3,
				y: 7,
			},
			want: 4,
		},
		{
			name: "calculator",
			args: args{
				x: 9,
				y: 3,
			},
			want: 6,
		},
		{
			name: "calculator",
			args: args{
				x: 4,
				y: 16,
			},
			want: 2,
		},
		{
			name: "calculator",
			args: args{
				x: 15,
				y: 3,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, brokenCalc(tt.args.x, tt.args.y))
		})
	}
}
