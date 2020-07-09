package Lowest_Common_Ancestor_Of_Deepest_Leaves_1123

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a ...int) int {
	res := a[0]
	for i := range a {
		if a[i] > res {
			res = a[i]
		}
	}
	return res
}

func postOrder(root, res *TreeNode, cnt int, maxlvl *int) int {
	if root == nil {
		return -1
	}
	if cnt > *maxlvl {
		*maxlvl = cnt
		*res = *root
	}
	ls := postOrder(root.Left, res, cnt+1, maxlvl)
	rs := postOrder(root.Right, res, cnt+1, maxlvl)
	if ls == rs && ls >= *maxlvl {
		*res = *root
		*maxlvl = ls
	}
	return max(ls, rs, cnt)
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	res := TreeNode{}
	maxlvl := -1
	postOrder(root, &res, 0, &maxlvl)
	return &res
}
