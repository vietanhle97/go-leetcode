package Unique_Binary_Search_Trees_II_95

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(start, end int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if start > end {
		res = append(res, nil)
		return res
	}
	for i := start; i <= end; i++ {
		ls := buildTree(start, i-1)
		rs := buildTree(i+1, end)
		for j := 0; j < len(ls); j++ {
			l := ls[j]
			for k := 0; k < len(rs); k++ {
				r := rs[k]
				root := &TreeNode{i, l, r}
				res = append(res, root)

			}
		}
	}

	return res
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	res := buildTree(1, n)
	return res
}
