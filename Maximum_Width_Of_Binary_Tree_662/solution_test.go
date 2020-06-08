package Maximum_Width_Of_Binary_Tree_662

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_widthOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "width of binary tree",
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
			want: 2,
		},
		{
			name: "width of binary tree",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 1,
							Left: &TreeNode{
								Val: 1,
							},
						},
					},
					Right: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val: 1,
							Right: &TreeNode{
								Val: 1,
							},
						},
					},
				},
			},
			want: 8,
		},
		{
			name: "width of binary tree",
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
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, widthOfBinaryTree(tt.args.root))
		})
	}
}
