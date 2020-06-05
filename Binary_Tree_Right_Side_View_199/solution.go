package Binary_Tree_Right_Side_View_199

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

func findOrder(node *TreeNode, lvl *Level, order int) {
	if node == nil {
		return
	}
	lvl.order = max(lvl.order, order)
	findOrder(node.Left, lvl, order+1)
	if _, ok := lvl.m[order]; ok {
		lvl.m[order] = append(lvl.m[order], node.Val)
	} else {
		lvl.m[order] = []int{node.Val}
	}
	findOrder(node.Right, lvl, order+1)
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	lvl := Level{map[int][]int{}, 0}
	findOrder(root, &lvl, 0)
	res := make([]int, lvl.order+1)
	i := 0
	for i <= lvl.order {
		res[i] = lvl.m[i][len(lvl.m[i])-1]
		i++
	}
	return res
}
