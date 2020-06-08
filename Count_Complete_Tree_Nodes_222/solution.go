package Count_Complete_Tree_Nodes_222

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(root *TreeNode, res *int) {
	if root == nil {
		return
	}
	*res += 1
	preOrder(root.Left, res)
	preOrder(root.Right, res)
}
func countNodes(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	preOrder(root, &res)
	return res
}
