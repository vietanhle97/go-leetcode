package Maximum_Sum_BST_In_Binary_Tree_1373

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func postOrder(root *TreeNode, cur []int, isBST *map[*TreeNode]bool, res *int) {
	if root == nil {
		return
	}

	// Calculate the sum of subtree
	ls, rs := []int{0}, []int{0}
	postOrder(root.Left, ls, isBST, res)
	postOrder(root.Right, rs, isBST, res)

	// sum of the current subtree
	cur[0] = ls[0] + rs[0] + root.Val

	// Check if the current root is the root of the valid BST
	if root.Left == nil && root.Right == nil {
		(*isBST)[root] = true
	} else if root.Left == nil {
		if root.Val < root.Right.Val {
			(*isBST)[root] = (*isBST)[root.Right]
		} else {
			(*isBST)[root] = false
		}
	} else if root.Right == nil {
		if root.Val > root.Left.Val {
			(*isBST)[root] = (*isBST)[root.Left]
		} else {
			(*isBST)[root] = false
		}
	} else {
		if root.Val > root.Left.Val && root.Val < root.Right.Val {
			(*isBST)[root] = (*isBST)[root.Left] && (*isBST)[root.Right]
		} else {
			(*isBST)[root] = false
		}
	}

	// If the root is a valid BST => update the return value
	// to the max of current value, the sum of left subtree, the sum of the right subtree
	// and the sum of the whole tree
	if (*isBST)[root] {
		*res = max(*res, max(max(ls[0], rs[0]), cur[0]))
	}

}

func maxSumBST(root *TreeNode) int {
	cur := []int{0}
	res := 0
	isBST := map[*TreeNode]bool{}
	postOrder(root, cur, &isBST, &res)
	return res
}
