package Binary_Tree_Coloring_Game_1145

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_binaryTreeColorGame(t *testing.T) {
	type args struct {
		root *TreeNode
		n    int
		x    int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "binary tree color game",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 8,
							},
							Right: &TreeNode{
								Val: 9,
							},
						},
						Right: &TreeNode{
							Val: 5,
							Left: &TreeNode{
								Val: 10,
							},
							Right: &TreeNode{
								Val: 11,
							},
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				n: 11,
				x: 3,
			},
			want: true,
		},
		{
			name: "binary tree color game",
			args: args{
				root: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 4,
							Right: &TreeNode{
								Val: 2,
								Right: &TreeNode{
									Val: 1,
									Right: &TreeNode{
										Val: 5,
									},
								},
							},
						},
					},
				},
				n: 7,
				x: 3,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, btreeGameWinningMove(tt.args.root, tt.args.n, tt.args.x))
		})
	}
}
