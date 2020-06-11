package Course_Schedule_II_210

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_courseScheduleII(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "course schedule II",
			args: args{
				numCourses:    5,
				prerequisites: [][]int{{1, 0}, {0, 2}, {4, 0}},
			},
			want: []int{2, 0, 1, 4, 3},
		},
		{
			name: "course schedule II",
			args: args{
				numCourses:    3,
				prerequisites: [][]int{{1, 0}},
			},
			want: []int{0, 1, 2},
		},
		{
			name: "course schedule II",
			args: args{
				numCourses:    4,
				prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			},
			want: []int{0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findOrder(tt.args.numCourses, tt.args.prerequisites))
		})
	}
}
