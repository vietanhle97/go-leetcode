package Distribute_Coins_In_Binary_Tree_979

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_distributeCoin(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "distribute coins in BT",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 0,
					},
				},
			},
			want: 2,
		},
		{
			name: "distribute coins in BT",
			args: args{
				root: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 0,
					},
				},
			},
			want: 3,
		},
		{
			name: "distribute coins in BT",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, distributeCoins(tt.args.root))
		})
	}
}
