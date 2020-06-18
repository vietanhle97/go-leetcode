package Basic_Calculator_II_227

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_calculator(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic calculator",
			args: args{
				s: "14 / 3 * 2",
			},
			want: 8,
		},
		{
			name: "basic calculator",
			args: args{
				s: " 3 + 7*9+6-421 ",
			},
			want: -349,
		},
		{
			name: "basic calculator",
			args: args{
				s: " -3 * 7-9+6-421 ",
			},
			want: -445,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, calculate(tt.args.s))
		})
	}
}
