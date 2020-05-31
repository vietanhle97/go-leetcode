package Binary_Tree_Preorder_Traversal_144

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_preOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "pre order",
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
			want: []int{1, 2, 3},
		},
		{
			name: "pre order",
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
			want: []int{1, -2, 3},
		},
		{
			name: "pre order",
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
			want: []int{-10, 9, 20, 15, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, preorderTraversal(tt.args.root))
		})
	}
}
