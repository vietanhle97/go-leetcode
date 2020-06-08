package Count_Good_Nodes_In_Binary_Tree_1448

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

func preOrder(root *TreeNode, m int, res *int) {
	if root == nil {
		return
	}
	if root.Val >= m {
		*res += 1
	}
	preOrder(root.Left, max(m, root.Val), res)
	preOrder(root.Right, max(m, root.Val), res)
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	preOrder(root, root.Val, &res)
	return res
}
