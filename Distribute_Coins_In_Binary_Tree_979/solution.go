package Distribute_Coins_In_Binary_Tree_979

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DFS(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	ls := DFS(root.Left, res)
	rs := DFS(root.Right, res)
	*res += int(math.Abs(float64(ls)) + math.Abs(float64(rs)))
	return root.Val + ls + rs - 1
}

func distributeCoins(root *TreeNode) int {
	res := 0
	DFS(root, &res)
	return res
}
