package Lowest_Common_Ancestor_Of_A_Binary_Tree_236

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Map struct {
	m map[int][]*TreeNode
}

func preOrder(root *TreeNode, m *Map) {
	if root == nil {
		return
	}
	if root.Left != nil {
		m.m[root.Left.Val] = []*TreeNode{root, root.Left}
	}
	if root.Right != nil {
		m.m[root.Right.Val] = []*TreeNode{root, root.Right}
	}
	preOrder(root.Left, m)
	preOrder(root.Right, m)
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	m := Map{map[int][]*TreeNode{}}
	m.m[root.Val] = []*TreeNode{root, root}
	preOrder(root, &m)

	check := map[int]bool{}
	curP := p.Val
	curQ := q.Val

	for curP != root.Val {
		check[curP] = true
		curP = m.m[curP][0].Val
	}
	check[root.Val] = true
	for curQ != root.Val {
		if _, ok := check[curQ]; ok {
			if curQ == q.Val {
				return q
			}
			return m.m[curQ][1]
		}
		curQ = m.m[curQ][0].Val
	}
	return root
}
