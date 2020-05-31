package Binary_Tree_Inorder_Traversal_94

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrder(root *TreeNode, output *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, output)
	*output = append(*output, root.Val)
	inOrder(root.Right, output)
}
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	inOrder(root, &res)
	return res
}
