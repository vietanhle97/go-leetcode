package Sum_Root_To_Leaf_Numbers_129

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(root *TreeNode, cur int, res *int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*res += cur*10 + root.Val
	}
	preOrder(root.Left, cur*10+root.Val, res)
	preOrder(root.Right, cur*10+root.Val, res)
}

func sumNumbers(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	preOrder(root, 0, &res)
	return res
}
