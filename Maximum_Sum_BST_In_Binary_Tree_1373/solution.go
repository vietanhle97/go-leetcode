package Maximum_Sum_BST_In_Binary_Tree_1373

import "math"

// I use the value 50000 and -50000 because -40000 <= node value <= 40000

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func DFS(root *TreeNode, res *int) (int, int, int, bool) {
	if root == nil {
		return 50000, -50000, 0, true
	}
	if root.Left == nil && root.Right == nil {
		*res = max(root.Val, *res)
		return root.Val, root.Val, root.Val, true
	}

	minLS, maxLS, sumLS, checkLS := DFS(root.Left, res)
	minRS, maxRS, sumRS, checkRS := DFS(root.Right, res)

	if checkLS && checkRS {
		if root.Val > maxLS && root.Val < minRS {
			cur := sumLS + sumRS + root.Val
			*res = max(*res, cur)
			return min(root.Val, minLS), max(root.Val, maxRS), cur, true
		}
	}
	return -50000, 50000, 0, false
}

func maxSumBST(root *TreeNode) int {
	res := 0
	DFS(root, &res)
	return res
}
