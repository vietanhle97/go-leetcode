package Construct_Binary_Tree_From_Inorder_And_Postorder_Traversal_106

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

func build(postOrder *[]int, inOrder *[]int) *TreeNode {
	if len(*inOrder) > 0 {
		root := &TreeNode{}
		m := len(*postOrder) - 1
		root.Val = (*postOrder)[m]
		ind := find(*inOrder, (*postOrder)[m])
		*postOrder = (*postOrder)[:m]
		if ind != -1 {
			l := (*inOrder)[:ind]
			r := (*inOrder)[ind+1:]
			root.Right = build(postOrder, &r)
			root.Left = build(postOrder, &l)
		}
		return root
	}
	return nil
}

func buildTree(inOrder []int, postOrder []int) *TreeNode {
	return build(&postOrder, &inOrder)
}
