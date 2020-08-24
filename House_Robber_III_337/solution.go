package House_Robber_III_337

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	l, r := dfs(root.Left), dfs(root.Right)
	return []int{max(l[0], l[1]) + max(r[0], r[1]), root.Val + l[0] + r[0]}

}

func rob(root *TreeNode) int {
	res := dfs(root)
	return max(res[0], res[1])
}
