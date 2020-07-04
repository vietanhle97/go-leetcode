package Maximum_Product_Of_Splitted_Binary_Tree_1339

import "math"

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

func getSumOfTreeNode(root *TreeNode, res *int) {
	if root == nil {
		return
	}
	getSumOfTreeNode(root.Left, res)
	getSumOfTreeNode(root.Right, res)
	*res += root.Val
}

func findMaxProductSubtree(root *TreeNode, cur []int, res *int, total int) {
	if root == nil {
		return
	}
	left, right := []int{0}, []int{0}
	findMaxProductSubtree(root.Left, left, res, total)
	findMaxProductSubtree(root.Right, right, res, total)
	cur[0] = left[0] + right[0] + root.Val
	*res = max(*res, (total-cur[0])*cur[0])
}

func maxProduct(root *TreeNode) int {
	cur := []int{0}
	total := 0
	res := 0
	getSumOfTreeNode(root, &total)
	findMaxProductSubtree(root, cur, &res, total)
	return res % (int(math.Pow10(9) + 7))
}
