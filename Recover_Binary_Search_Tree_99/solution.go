package Recover_Binary_Search_Tree_99

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrder(root *TreeNode, prev **TreeNode, order **[][]*TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left, prev, order)
	if *prev != nil && (*prev).Val > root.Val {
		**order = append(**order, []*TreeNode{*prev, root})
	}
	*prev = root
	inOrder(root.Right, prev, order)
}

func swap(r1, r2 *TreeNode) {
	r1.Val, r2.Val = r2.Val, r1.Val
}

func recoverTree(root *TreeNode) {
	order := make([][]*TreeNode, 0)
	tmp := &order
	var prev *TreeNode
	inOrder(root, &prev, &(tmp))
	if len(order) == 2 {
		swap(order[0][0], order[1][1])
	} else if len(order) == 1 {
		swap(order[0][0], order[0][1])
	}
}
