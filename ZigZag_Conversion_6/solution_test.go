package ZigZag_Conversion_6

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_zigzagConversion(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "zigzag conversion",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 2,
			},
			want: "PYAIHRNAPLSIIG",
		},
		{
			name: "zigzag conversion",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 3,
			},
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "zigzag conversion",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 4,
			},
			want: "PINALSIGYAHRPI",
		},
		{
			name: "zigzag conversion",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 5,
			},
			want: "PHASIYIRPLIGAN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, convert(tt.args.s, tt.args.numRows))
		})
	}
}
