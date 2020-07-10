package Flatten_A_Multilevel_Doubly_Linked_List_430

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func preOrder(root *Node, m *[]*Node) {
	if root == nil {
		return
	}
	*m = append(*m, root)
	preOrder(root.Child, m)
	preOrder(root.Next, m)

}

func flatten(root *Node) *Node {
	m := make([]*Node, 0)
	preOrder(root, &m)
	for i := 0; i < len(m)-1; i++ {
		m[i].Next = m[i+1]
		m[i+1].Prev = m[i]
		m[i].Child = nil
		m[i+1].Child = nil
	}

	return root
}
