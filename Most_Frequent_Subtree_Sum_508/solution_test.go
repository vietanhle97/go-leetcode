package Most_Frequent_Subtree_Sum_508

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_mostFreqSubtreeSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "binary tree maximum path sum",
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
			want: []int{2, 3, 6},
		},
		{
			name: "binary tree maximum path sum",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: -5,
					},
				},
			},
			want: []int{2},
		},
		{
			name: "binary tree maximum path sum",
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
			want: []int{9, 15, 7, 42, 41},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Ints(tt.want)
			res := findFrequentTreeSum(tt.args.root)
			sort.Ints(res)
			assert.Equal(t, tt.want, res)
		})
	}
}
