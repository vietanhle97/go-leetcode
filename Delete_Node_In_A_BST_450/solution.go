package Delete_Node_In_A_BST_450

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if key == root.Val {
		r, l := root.Right, root.Left
		if l == nil && r == nil {
			return nil
		}
		if l == nil {
			return r
		}
		if r == nil {
			return l
		}
		for r != nil && r.Left != nil {
			r = r.Left
		}
		r.Left = root.Left
		return root.Right
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}
