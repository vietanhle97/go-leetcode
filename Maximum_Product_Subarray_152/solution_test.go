package Maximum_Product_Subarray_152

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxProductSub(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "maximum subarray product",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 24,
		},
		{
			name: "maximum subarray product",
			args: args{
				nums: []int{-2, 2, 3, -4},
			},
			want: 48,
		},
		{
			name: "maximum subarray product",
			args: args{
				nums: []int{-2, 0, -1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxProduct(tt.args.nums))
		})
	}
}
