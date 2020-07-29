package Flip_Binary_Tree_To_Match_Preorder_Traversal_971

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, voyage []int, res *[]int, cnt *int) {
	if root == nil {
		return
	}
	if root.Val != voyage[*cnt] {
		*res = []int{-1}
		return
	}
	*cnt++
	if *cnt < len(voyage) && root.Left != nil && root.Left.Val != voyage[*cnt] {
		*res = append(*res, root.Val)
		dfs(root.Right, voyage, res, cnt)
		dfs(root.Left, voyage, res, cnt)
	} else {
		dfs(root.Left, voyage, res, cnt)
		dfs(root.Right, voyage, res, cnt)
	}
}

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	res, cnt := make([]int, 0), 0
	dfs(root, voyage, &res, &cnt)
	return res
}
