package Find_Duplicate_Subtrees_652

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postOrder(root *TreeNode, ind string, cur []string, vis map[string]bool, m map[string]*TreeNode, res *[]*TreeNode) {
	if root == nil {
		cur[0] = ind + "#"
		return
	}
	ls, rs := []string{""}, []string{""}
	postOrder(root.Left, "l", ls, vis, m, res)
	postOrder(root.Right, "r", rs, vis, m, res)
	cur[0] = ls[0] + strconv.Itoa(root.Val) + rs[0]
	if _, ok := m[cur[0]]; !ok {
		m[cur[0]] = root
		vis[cur[0]] = false
	} else {
		if _, ok := vis[cur[0]]; ok && !vis[cur[0]] {
			*res = append(*res, root)
			vis[cur[0]] = true
		}
	}

}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	m, vis, res, cur := map[string]*TreeNode{}, map[string]bool{}, make([]*TreeNode, 0), []string{""}
	postOrder(root, "", cur, vis, m, &res)
	return res
}
