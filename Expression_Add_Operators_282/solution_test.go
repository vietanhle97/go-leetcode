package Expression_Add_Operators_282

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_expressAddOperators(t *testing.T) {
	type args struct {
		num    string
		target int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "expression add operators",
			args: args{
				num:    "105",
				target: 5,
			},
			want: []string{"1*0+5", "10-5"},
		},
		{
			name: "expression add operators",
			args: args{
				num:    "123",
				target: 6,
			},
			want: []string{"1+2+3", "1*2*3"},
		},
		{
			name: "expression add operators",
			args: args{
				num:    "232",
				target: 8,
			},
			want: []string{"2+3*2", "2*3+2"},
		},
		{
			name: "expression add operators",
			args: args{
				num:    "00",
				target: 0,
			},
			want: []string{"0+0", "0*0", "0-0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, addOperators(tt.args.num, tt.args.target))
		})
	}
}
