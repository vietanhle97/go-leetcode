package Basic_Calculator_224

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
				s: "(1+(4+5+2)-3)+(6+8)",
			},
			want: 23,
		},
		{
			name: "basic calculator",
			args: args{
				s: " 3 + 7-9+6-421 ",
			},
			want: -414,
		},
		{
			name: "basic calculator",
			args: args{
				s: " -3 + 7-9+6-421 ",
			},
			want: -420,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, calculate(tt.args.s))
		})
	}
}
