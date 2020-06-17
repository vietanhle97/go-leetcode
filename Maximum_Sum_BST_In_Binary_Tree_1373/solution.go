package Maximum_Sum_BST_In_Binary_Tree_1373

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

func DFS(root *TreeNode, maxSum *int) (int, int, int, bool) {
	if root == nil {
		return 50000, -50000, 0, true
	}
	if root.Left == nil && root.Right == nil {
		*maxSum = max(*maxSum, root.Val)
		return root.Val, root.Val, root.Val, true
	}

	minLS, maxLS, sumLS, checkLS := DFS(root.Left, maxSum)
	minRS, maxRS, sumRS, checkRS := DFS(root.Right, maxSum)

	if checkRS && checkLS {
		if root.Val > maxLS && root.Val < minRS {
			cur := sumLS + sumRS + root.Val
			*maxSum = max(*maxSum, cur)
			return min(root.Val, minLS), max(root.Val, maxRS), cur, true
		}
	}

	return -50000, 50000, *maxSum, false
}

func maxSumBST(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	DFS(root, &res)
	return res
}
