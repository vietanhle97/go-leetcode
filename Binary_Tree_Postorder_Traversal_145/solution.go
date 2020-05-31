package Binary_Tree_Postorder_Traversal_145

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postOrder(root *TreeNode, output *[]int) {
	if root == nil {
		return
	}
	postOrder(root.Left, output)
	postOrder(root.Right, output)
	*output = append(*output, root.Val)
}
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	postOrder(root, &res)
	return res
}
