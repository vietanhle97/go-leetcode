package Maximum_Sum_BST_In_Binary_Tree_1373

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxSumBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "max sum BST in Binary Tree",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
			},
			want: 2,
		},
		{
			name: "max sum BST in Binary Tree",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
				},
			},
			want: 7,
		},
		{
			name: "max sum BST in Binary Tree",
			args: args{
				root: &TreeNode{
					Val: -4,
					Left: &TreeNode{
						Val: -2,
					},
					Right: &TreeNode{
						Val: -5,
					},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxSumBST(tt.args.root))
		})
	}
}
