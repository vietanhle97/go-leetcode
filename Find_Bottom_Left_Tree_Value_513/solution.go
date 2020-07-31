package Find_Bottom_Left_Tree_Value_513

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postOrder(root *TreeNode, level, ind int, lv, i, res *int) {
	if root == nil {
		return
	}
	postOrder(root.Left, level+1, 2*ind, lv, i, res)
	postOrder(root.Right, level+1, 2*ind+1, lv, i, res)
	if *lv < level {
		*lv, *i, *res = level, ind, root.Val
	} else if *lv == level && *i > ind {
		*lv, *i, *res = level, ind, root.Val
	}
}

func findBottomLeftValue(root *TreeNode) int {
	lv, i, res := -1, -1, -1
	postOrder(root, 0, 1, &lv, &i, &res)
	return res
}
