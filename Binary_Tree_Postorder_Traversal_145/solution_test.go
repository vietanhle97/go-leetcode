package Binary_Tree_Postorder_Traversal_145

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_postOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "post order",
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
			want: []int{2, 3, 1},
		},
		{
			name: "post order",
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
			want: []int{-2, 3, 1},
		},
		{
			name: "post order",
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
			want: []int{9, 15, 7, 20, -10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, postorderTraversal(tt.args.root))
		})
	}
}
