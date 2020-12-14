package Convert_BST_To_Greater_Tree_538

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs1(root *TreeNode, cur []int) {
	if root == nil {
		return
	}
	l, r := []int{0}, []int{0}
	dfs1(root.Left, l)
	dfs1(root.Right, r)
	cur[0] = l[0] + r[0] + root.Val
	root.Val = root.Val + r[0]
}

func dfs2(root *TreeNode, val int) {
	if root == nil {
		return
	}
	tmp := root.Val
	root.Val = root.Val + val
	dfs2(root.Left, root.Val)
	dfs2(root.Right, root.Val-tmp)

}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	dfs1(root, []int{0})
	dfs2(root.Left, root.Val)
	tmp := root.Right
	for tmp != nil {
		dfs2(tmp.Left, tmp.Val)
		tmp = tmp.Right
	}
	return root
}
