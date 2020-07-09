package Sum_Of_Nodes_With_Even_Valued_Grandparent_1315

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(gp, p int, root *TreeNode, res *int) {
	if root == nil {
		return
	}
	if gp != -1 && gp%2 == 0 {
		*res += root.Val
	}
	if p == -1 {
		preOrder(-1, root.Val, root.Left, res)
		preOrder(-1, root.Val, root.Right, res)
	} else {
		preOrder(p, root.Val, root.Left, res)
		preOrder(p, root.Val, root.Right, res)
	}
}

func sumEvenGrandparent(root *TreeNode) int {
	res := 0
	preOrder(-1, -1, root, &res)
	return res
}
