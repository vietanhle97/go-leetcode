package Construct_Binary_Search_Tree_From_Preorder_Traversal_1008

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func find(nums []int, val int) int {
	for i := range nums {
		if nums[i] > val {
			return i
		}
	}
	return -1
}

func bstFromPreorder(preOrder []int) *TreeNode {
	if len(preOrder) == 0 {
		return nil
	}
	val := preOrder[0]
	root := &TreeNode{val, nil, nil}
	ind := find(preOrder, val)
	if ind != -1 {
		root.Left = bstFromPreorder(preOrder[1:ind])
		root.Right = bstFromPreorder(preOrder[ind:])
	} else {
		root.Left = bstFromPreorder(preOrder[1:])
	}
	return root
}
