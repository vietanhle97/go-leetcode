package Binary_Tree_Pruning_814

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs1(root *TreeNode, cur []int, m map[*TreeNode]bool) {
	if root == nil {
		return
	}
	l, r := []int{0}, []int{0}
	dfs1(root.Left, l, m)
	dfs1(root.Right, r, m)
	if root.Val == 1 {
		cur[0] = l[0] + r[0] + 1
	} else {
		cur[0] = l[0] + r[0]
	}
	if cur[0] == 0 {
		m[root] = true
	}
}

func dfs2(root *TreeNode, cnt int, m map[*TreeNode]bool) {
	if root == nil {
		return
	}
	if _, ok := m[root.Left]; ok && root.Left != nil {
		root.Left = nil
	}
	if _, ok := m[root.Right]; ok && root.Right != nil {
		root.Right = nil
	}
	if root.Left != nil {
		dfs2(root.Left, 2*cnt, m)
	}
	if root.Right != nil {
		dfs2(root.Right, 2*cnt+1, m)
	}
}

func pruneTree(root *TreeNode) *TreeNode {
	cur, m := []int{0}, map[*TreeNode]bool{}
	dfs1(root, cur, m)
	if _, ok := m[root]; ok {
		return nil
	}
	dfs2(root, 1, m)
	return root
}
