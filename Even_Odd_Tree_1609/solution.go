package Even_Odd_Tree_1609

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, level int, m map[int]int, res *bool) {
	if root == nil || !(*res) {
		return
	}
	if level > 0 {
		if level%2 == 0 {
			if _, ok := m[level]; ok && m[level] >= root.Val || root.Val%2 == 0 {
				*res = false
				return
			}
			m[level] = root.Val
		} else {
			if _, ok := m[level]; ok && m[level] <= root.Val || root.Val%2 == 1 {
				*res = false
				return
			}
			m[level] = root.Val
		}
	} else {
		if root.Val%2 == 0 {
			*res = false
			return
		}
	}
	dfs(root.Left, level+1, m, res)
	dfs(root.Right, level+1, m, res)
}

func isEvenOddTree(root *TreeNode) bool {
	m := map[int]int{}
	res := true
	dfs(root, 0, m, &res)
	return res
}
