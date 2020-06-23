package Sum_Of_Mutated_Array_Closest_To_Target_1300

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sumClosestToTarget(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sum of mutated array closest to target",
			args: args{
				arr:    []int{3, 4, 9},
				target: 10,
			},
			want: 3,
		},
		{
			name: "sum of mutated array closest to target",
			args: args{
				arr:    []int{3, 4, 9},
				target: 17,
			},
			want: 9,
		},
		{
			name: "sum of mutated array closest to target",
			args: args{
				arr:    []int{1, 2, 3, 20, 24, 34, 36},
				target: 110,
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findBestValue(tt.args.arr, tt.args.target))
		})
	}
}
