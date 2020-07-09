package Maximum_Binary_Tree_654

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMax(nums []int) int {
	ind := 0
	for i := range nums {
		if nums[i] > nums[ind] {
			ind = i
		}
	}
	return ind
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max := findMax(nums)
	root := &TreeNode{nums[max], nil, nil}
	root.Left = constructMaximumBinaryTree(nums[:max])
	root.Right = constructMaximumBinaryTree(nums[max+1:])
	return root
}
