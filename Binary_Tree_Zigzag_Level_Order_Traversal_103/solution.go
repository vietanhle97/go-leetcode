package Binary_Tree_Zigzag_Level_Order_Traversal_103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Level struct {
	m     map[int][]int
	order int
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func reverse(m []int) []int {
	n := len(m)
	res := make([]int, n)
	for i, _ := range m {
		res[i] = m[n-1-i]
	}
	return res
}

func zigzagOrder(node *TreeNode, lvl *Level, order int) {
	if node == nil {
		return
	}
	lvl.order = max(lvl.order, order)
	zigzagOrder(node.Left, lvl, order+1)
	if _, ok := lvl.m[order]; ok {
		lvl.m[order] = append(lvl.m[order], node.Val)
	} else {
		lvl.m[order] = []int{node.Val}
	}
	zigzagOrder(node.Right, lvl, order+1)
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	lvl := Level{map[int][]int{}, 0}
	zigzagOrder(root, &lvl, 0)
	res := make([][]int, lvl.order+1)
	i := 0
	for i <= lvl.order {
		if i%2 == 1 {
			res[i] = reverse(lvl.m[i])
		} else {
			res[i] = lvl.m[i]
		}
		i++
	}
	return res
}
