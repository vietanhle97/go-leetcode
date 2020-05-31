package Course_Schdule_207

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_courseSchedule(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "course schedule",
			args: args{
				prerequisites: [][]int{{1, 0}},
				numCourses:    2,
			},
			want: true,
		},
		{
			name: "course schedule",
			args: args{
				prerequisites: [][]int{{1, 0}, {0, 1}},
				numCourses:    2,
			},
			want: false,
		},
		{
			name: "course schedule",
			args: args{
				prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}},
				numCourses:    4,
			},
			want: true,
		},
		{
			name: "course schedule",
			args: args{
				prerequisites: [][]int{{1, 0}, {2, 0}, {0, 3}, {3, 2}},
				numCourses:    4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, canFinish(tt.args.numCourses, tt.args.prerequisites))
		})
	}
}
