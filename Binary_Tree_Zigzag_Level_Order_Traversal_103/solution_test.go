package Binary_Tree_Zigzag_Level_Order_Traversal_103

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_zigzagOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "zigzag order",
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
			want: [][]int{{1}, {3, 2}},
		},
		{
			name: "zigzag order",
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
			want: [][]int{{1}, {3, -2}},
		},
		{
			name: "zigzag order",
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
			want: [][]int{{-10}, {20, 9}, {15, 7}},
		},
		{
			name: "zigzag order",
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
			want: [][]int{{1}, {3, 2}, {4, 5, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, zigzagLevelOrder(tt.args.root))
		})
	}
}
