package Maximum_Binary_Tree_II_998

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

func inOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, res)
	*res = append(*res, root.Val)
	inOrder(root.Right, res)
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	res := make([]int, 0)
	inOrder(root, &res)
	res = append(res, val)
	return constructMaximumBinaryTree(res)
}
