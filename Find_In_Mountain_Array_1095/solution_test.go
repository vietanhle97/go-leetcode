package Find_In_Mountain_Array_1095

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_findInMountainArray(t *testing.T) {
	type args struct {
		mountainArray *MountainArray
		target        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find in mountain array",
			args: args{
				mountainArray: &MountainArray{
					nums: []int{1, 2, 3, 4, 5, 3, 1},
				},
				target: 3,
			},
			want: 2,
		},
		{
			name: "find in mountain array",
			args: args{
				mountainArray: &MountainArray{
					nums: []int{0, 1, 2, 4, 2, 1},
				},
				target: 3,
			},
			want: -1,
		},
		{
			name: "find in mountain array",
			args: args{
				mountainArray: &MountainArray{
					nums: []int{1, 2, 3, 4, 5, 3, 1},
				},
				target: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findInMountainArray(tt.args.target, tt.args.mountainArray))
		})
	}
}
