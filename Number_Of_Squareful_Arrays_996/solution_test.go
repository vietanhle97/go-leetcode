package Number_Of_Squareful_Arrays_996

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numOfSquarefulArr(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "number of squareful arrays",
			args: args{
				nums: []int{1, 17, 8},
			},
			want: 2,
		},
		{
			name: "number of squareful arrays",
			args: args{
				nums: []int{5, 11, 5, 4, 5},
			},
			want: 2,
		},
		{
			name: "number of squareful arrays",
			args: args{
				nums: []int{2, 2, 2},
			},
			want: 1,
		},
		{
			name: "number of squareful arrays",
			args: args{
				nums: []int{1, 1, 8, 1, 8},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, numSquarefulPerms(tt.args.nums))
		})
	}
}
