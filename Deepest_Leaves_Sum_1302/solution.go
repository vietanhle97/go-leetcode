package Deepest_Leaves_Sum_1302

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
func postOrder(root *TreeNode, res, MAX *int, level int) {
	if root == nil {
		return
	}
	postOrder(root.Left, res, MAX, level+1)
	postOrder(root.Right, res, MAX, level+1)
	if root.Left == nil && root.Right == nil {
		if level > *MAX {
			*res = root.Val
		} else if level == *MAX {
			*res += root.Val
		}
	}
	*MAX = max(*MAX, level)
}
func deepestLeavesSum(root *TreeNode) int {
	res := 0
	MAX := 0
	postOrder(root, &res, &MAX, 0)
	return res
}
