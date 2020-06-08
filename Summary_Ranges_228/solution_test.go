package Summary_Ranges_228

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "summary ranges",
			args: args{
				nums: []int{0, 2, 3, 4, 6, 8, 9},
			},
			want: []string{"0", "2->4", "6", "8->9"},
		},
		{
			name: "summary ranges",
			args: args{
				nums: []int{0, 1, 2, 4, 5, 6, 7},
			},
			want: []string{"0->2", "4->7"},
		},
		{
			name: "summary ranges",
			args: args{
				nums: []int{0, 1, 2, 4, 5, 7},
			},
			want: []string{"0->2", "4->5", "7"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, summaryRanges(tt.args.nums))
		})
	}
}
