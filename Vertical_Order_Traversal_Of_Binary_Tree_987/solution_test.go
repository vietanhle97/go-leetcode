package Vertical_Order_Traversal_Of_Binary_Tree_987

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_verticalOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "vertical order",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: [][]int{{2}, {1}, {3}},
		},
		{
			name: "vertical order",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: -2,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: [][]int{{-2}, {1}, {3}},
		},
		{
			name: "vertical order",
			args: args{
				root: &TreeNode{
					Val: -10,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 20,
						Left: &TreeNode{
							Val: 15,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			want: [][]int{{9}, {-10}, {15}, {20}, {7}},
		},
		{
			name: "vertical order",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
			want: [][]int{{4}, {2}, {1}, {5}, {3}, {6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, verticalTraversal(tt.args.root))
		})
	}
}
