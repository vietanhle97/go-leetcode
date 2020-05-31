package Binary_Tree_Maximum_Path_Sum_124

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxOfThree(a int, b int, c int) int {
	if a >= b && b >= c {
		return a
	} else if b >= a && b >= c {
		return b
	} else if c >= a && c >= b {
		return c
	} else {
		return a
	}
}

func DFS(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	l, r := 0, 0
	ls, rs := -99999999999, -99999999999
	if node.Left != nil {
		l, ls = DFS(node.Left)
		l = max(l, 0)
	}
	if node.Right != nil {
		r, rs = DFS(node.Right)
		r = max(r, 0)
	}
	return node.Val + max(l, r), maxOfThree(node.Val+l+r, ls, rs)
}

func maxPathSum(root *TreeNode) int {
	return max(DFS(root))
}
