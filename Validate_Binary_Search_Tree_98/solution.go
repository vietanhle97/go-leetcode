package Validate_Binary_Search_Tree_98

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrder(root *TreeNode, output *[]int, res *bool) {
	if !(*res) {
		return
	}
	if root == nil {
		return
	}
	inOrder(root.Left, output, res)
	if len(*output) == 0 {
		*output = append(*output, root.Val)
	} else {
		if root.Val <= (*output)[len(*output)-1] {
			*res = false
			return
		} else {
			*output = append(*output, root.Val)
		}
	}

	inOrder(root.Right, output, res)
}

func isValidBST(root *TreeNode) bool {
	output := make([]int, 0)
	res := true
	inOrder(root, &output, &res)
	return res
}
