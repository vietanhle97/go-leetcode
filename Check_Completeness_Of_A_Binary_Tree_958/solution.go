package Check_Completeness_Of_A_Binary_Tree_958

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	node *TreeNode
	cnt  int
}

func isCompleteTree(root *TreeNode) bool {
	m := []*Node{&Node{root, 1}}
	i := 0
	for i < len(m) {
		cur := m[i]
		i++
		if cur.node != nil {
			m = append(m, &Node{cur.node.Left, 2 * cur.cnt})
			m = append(m, &Node{cur.node.Right, 2*cur.cnt + 1})
		}
	}
	return m[len(m)-1].cnt == len(m)
}
