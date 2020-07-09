package Maximum_Difference_Between_Node_And_Ancestor_1026_

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func preOrder(root *TreeNode, MIN, MAX int, res *int) {
	if root == nil {
		return
	}
	MIN = min(root.Val, MIN)
	MAX = max(root.Val, MAX)
	*res = max(MAX-MIN, *res)
	preOrder(root.Left, MIN, MAX, res)
	preOrder(root.Right, MIN, MAX, res)
}

func maxAncestorDiff(root *TreeNode) int {
	res := 0
	preOrder(root, root.Val, root.Val, &res)
	return res
}
