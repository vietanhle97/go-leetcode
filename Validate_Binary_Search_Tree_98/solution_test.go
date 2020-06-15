package Validate_Binary_Search_Tree_98

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_validatedBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "validated BST",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 1,
							Left: &TreeNode{
								Val: 0,
							},
							Right: &TreeNode{
								Val: 2,
							},
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 7,
							Left: &TreeNode{
								Val: 6,
							},
						},
						Right: &TreeNode{
							Val: 9,
							Right: &TreeNode{
								Val: 10,
							},
						},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isValidBST(tt.args.root))
		})
	}
}
