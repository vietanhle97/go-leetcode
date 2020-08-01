package Binary_Tree_Maximum_Path_Sum_124

const MinInt = -(1<<31 - 1) - 1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a ...int) int {
	res := a[0]
	for _, e := range a {
		if e > res {
			res = e
		}
	}
	return res
}

func dfs(root *TreeNode) (int, int) {
	if root == nil {
		return MinInt, MinInt
	}
	l, r, ls, rs := 0, 0, MinInt, MinInt
	if root.Left != nil {
		l, ls = dfs(root.Left)
		l = max(l, 0)
	}
	if root.Right != nil {
		r, rs = dfs(root.Right)
		r = max(r, 0)
	}
	return root.Val + max(l, r), max(ls, rs, root.Val+l+r)
}

func maxPathSum(root *TreeNode) int {
	return max(dfs(root))
}
