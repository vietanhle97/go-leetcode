package Binary_Search_Tree_To_Greater_Sum_Tree_1038

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postOrder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	postOrder(root.Left, nums)
	postOrder(root.Right, nums)
	*nums = append(*nums, root.Val)
}

func updateVal(root *TreeNode, newVal map[int]int) {
	if root == nil {
		return
	}
	root.Val = newVal[root.Val]
	updateVal(root.Left, newVal)
	updateVal(root.Right, newVal)
}

func bstToGst(root *TreeNode) *TreeNode {
	nums := make([]int, 0)
	postOrder(root, &nums)
	sort.Ints(nums)
	newVal := map[int]int{}
	sum := 0
	m := len(nums)
	for i := range nums {
		sum += nums[m-1-i]
		newVal[nums[m-1-i]] = sum
	}
	updateVal(root, newVal)
	return root
}
