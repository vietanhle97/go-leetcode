package Path_Sum_II_113

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "sum to target",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 11,
							Left: &TreeNode{
								Val: 7,
							},
							Right: &TreeNode{
								Val: 2,
							},
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 13,
						},
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 5,
							},
							Right: &TreeNode{
								Val: 1,
							},
						},
					},
				},
				target: 22,
			},
			want: [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			name: "sum to target",
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
				target: 4,
			},
			want: [][]int{{1, 3}},
		},
		{
			name: "sum to target",
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
				target: 25,
			},
			want: [][]int{{-10, 20, 15}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, pathSum(tt.args.root, tt.args.target))
		})
	}
}
