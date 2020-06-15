package Unique_Binary_Search_Trees_96

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_uniqueBST(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "unique BST",
			args: args{
				n: 7,
			},
			want: 429,
		},
		{
			name: "unique BST",
			args: args{
				n: 19,
			},
			want: 1767263190,
		},
		{
			name: "unique BST",
			args: args{
				n: 10,
			},
			want: 16796,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, numTrees(tt.args.n))
		})
	}
}
