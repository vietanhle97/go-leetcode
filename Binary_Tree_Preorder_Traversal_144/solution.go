package Binary_Tree_Preorder_Traversal_144

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(root *TreeNode, output *[]int) {
	if root == nil {
		return
	}
	*output = append(*output, root.Val)
	preOrder(root.Left, output)
	preOrder(root.Right, output)
}
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	preOrder(root, &res)
	return res
}
