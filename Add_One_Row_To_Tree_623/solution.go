package Add_One_Row_To_Tree_623

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(root *TreeNode, row, order, v int) {
	if root == nil || order > row {
		return
	}

	if row-1 == order {
		left := root.Left
		right := root.Right
		root.Left = &TreeNode{v, left, nil}
		root.Right = &TreeNode{v, nil, right}
	}
	preOrder(root.Left, row, order+1, v)
	preOrder(root.Right, row, order+1, v)
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		newRoot := &TreeNode{v, root, nil}
		return newRoot
	}
	preOrder(root, d, 1, v)
	return root
}
