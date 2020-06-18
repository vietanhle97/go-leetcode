package Largest_Number_179

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_largestNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "largest number",
			args: args{
				nums: []int{824, 938, 1399, 5607, 6973, 5703, 9609, 4398, 8247},
			},
			want: "9609938824824769735703560743981399",
		},
		{
			name: "largest number",
			args: args{
				nums: []int{0, 0},
			},
			want: "0",
		},
		{
			name: "largest number",
			args: args{
				nums: []int{121, 12},
			},
			want: "12121",
		},
		{
			name: "largest number",
			args: args{
				nums: []int{3, 30, 34, 5, 9},
			},
			want: "9534330",
		},
		{
			name: "largest number",
			args: args{
				nums: []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
			want: "9876543210",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, largestNumber(tt.args.nums))
		})
	}
}
