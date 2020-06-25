package Split_Array_Into_Fibonacci_Sequence_842

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_splitArrayToFibSequence(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "split array into fibonacci sequence",
			args: args{
				S: "123456579",
			},
			want: []int{123, 456, 579},
		},
		{
			name: "split array into fibonacci sequence",
			args: args{
				S: "11235813",
			},
			want: []int{1, 1, 2, 3, 5, 8, 13},
		},
		{
			name: "split array into fibonacci sequence",
			args: args{
				S: "112358130",
			},
			want: []int{},
		},
		{
			name: "split array into fibonacci sequence",
			args: args{
				S: "549132235579214924139063110211652267343256998",
			},
			want: []int{5, 4, 9, 13, 22, 35, 57, 92, 149, 241, 390, 631, 1021, 1652, 2673, 4325, 6998},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, splitIntoFibonacci(tt.args.S))
		})
	}
}
