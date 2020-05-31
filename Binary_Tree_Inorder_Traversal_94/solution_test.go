package Binary_Tree_Inorder_Traversal_94

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_inOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "in order",
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
			want: []int{2, 1, 3},
		},
		{
			name: "in order",
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
			want: []int{-2, 1, 3},
		},
		{
			name: "in order",
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
			want: []int{9, -10, 15, 20, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, inorderTraversal(tt.args.root))
		})
	}
}
