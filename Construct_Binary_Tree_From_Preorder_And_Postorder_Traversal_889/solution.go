package Construct_Binary_Tree_From_Preorder_And_Postorder_Traversal_889

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func build(pre, post *[]int) *TreeNode {
	m := len(*pre)
	if m == 0 {
		return nil
	}
	root := &TreeNode{}
	root.Val = (*pre)[0]
	if m == 1 {
		return root
	}

	l := 0
	for i := 0; i < m; i++ {
		if (*post)[i] == (*pre)[1] {
			l = i + 1
		}
	}
	preL := (*pre)[1 : l+1]
	postL := (*post)[:l]
	preR := (*pre)[l+1:]
	postR := (*post)[l : m-1]
	root.Left = build(&preL, &postL)
	root.Right = build(&preR, &postR)

	return root

}

func constructFromPrePost(pre []int, post []int) *TreeNode {
	return build(&pre, &post)
}
