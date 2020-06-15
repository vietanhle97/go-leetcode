package Balance_A_Binary_Search_Tree_1382

import (
	"fmt"
	"math"
)

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

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := height(root.Left), height(root.Right)
	if left == -1 || right == -1 || int(math.Abs(float64(left)-float64(right))) > 1 {
		return -1
	}
	return 1 + max(left, right)
}

func rotateLeft(root *TreeNode) *TreeNode {
	newRoot := root.Right
	root.Right = newRoot.Left
	newRoot.Left = root
	return newRoot
}

func rotateRight(root *TreeNode) *TreeNode {
	newRoot := root.Left
	root.Left = newRoot.Right
	newRoot.Right = root
	return newRoot
}

func balanceBST(root *TreeNode) *TreeNode {

	//for height(root) == -1 {
	//	l, r := height(root.Left), height(root.Right)
	//	if r > l {
	//		if height(root.Right.Left) > height(root.Right.Right) {
	//			root.Right = rotateRight(root.Right)
	//		}
	//		root = rotateLeft(root)
	//	} else {
	//		if height(root.Left.Right) > height(root.Left.Left) {
	//			root.Left = rotateLeft(root.Left)
	//		}
	//		root = rotateRight(root)
	//	}
	//}
	fmt.Println(height(root))
	return root
}
