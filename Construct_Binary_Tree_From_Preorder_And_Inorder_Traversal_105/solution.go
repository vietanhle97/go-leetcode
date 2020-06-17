package Construct_Binary_Tree_From_Preorder_And_Inorder_Traversal_105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func find(nums []int, v int) int {
	for i := range nums {
		if nums[i] == v {
			return i
		}
	}
	return -1
}

func build(preOrder *[]int, inOrder *[]int) *TreeNode {
	if len(*inOrder) > 0 {
		root := &TreeNode{}
		root.Val = (*preOrder)[0]
		ind := find(*inOrder, (*preOrder)[0])
		*preOrder = (*preOrder)[1:]
		if ind != -1 {
			l := (*inOrder)[:ind]
			r := (*inOrder)[ind+1:]
			root.Left = build(preOrder, &l)
			root.Right = build(preOrder, &r)
		}
		return root
	}
	return nil
}

func buildTree(preOrder []int, inOrder []int) *TreeNode {
	return build(&preOrder, &inOrder)
}
