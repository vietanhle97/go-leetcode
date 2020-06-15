package Validate_Binary_Search_Tree_98

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

func isValidBST(root *TreeNode) bool {
	res := make([]int, 0)
	inOrder(root, &res)
	for i := 1; i < len(res); i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}
	return true
}
