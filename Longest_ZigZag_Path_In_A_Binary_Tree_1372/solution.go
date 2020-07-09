package Longest_ZigZag_Path_In_A_Binary_Tree_1372

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func preOrder(root *TreeNode, cur, cnt int, res *int) {
	if root == nil {
		return
	}
	*res = max(*res, cnt)
	if cur == -1 {
		preOrder(root.Left, 0, cnt+1, res)
		preOrder(root.Right, 1, cnt+1, res)
	} else if cur == 0 {
		preOrder(root.Left, 0, 1, res)
		preOrder(root.Right, 1, cnt+1, res)
	} else {
		preOrder(root.Left, 0, cnt+1, res)
		preOrder(root.Right, 1, 1, res)
	}
}

func longestZigZag(root *TreeNode) int {
	res := 0
	preOrder(root, -1, 0, &res)
	return res
}
