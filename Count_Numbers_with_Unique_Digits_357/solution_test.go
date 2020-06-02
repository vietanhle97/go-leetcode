package Count_Numbers_with_Unique_Digits_357

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countNumbers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "count numbers",
			args: args{
				n: 3,
			},
			want: 739,
		},
		{
			name: "count numbers",
			args: args{
				n: 9,
			},
			want: 5611771,
		},
		{
			name: "count numbers",
			args: args{
				n: 2,
			},
			want: 91,
		},
		{
			name: "count numbers",
			args: args{
				n: 15,
			},
			want: 8877691,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countNumbersWithUniqueDigits(tt.args.n))
		})
	}
}
