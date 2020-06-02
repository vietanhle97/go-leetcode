package Populating_Next_Right_Pointers_In_Each_Node_116

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
type Level struct {
	m map[int][]*Node
}

func postOrder(node *Node, lvl *Level, order int) {
	if node == nil {
		return
	}
	postOrder(node.Left, lvl, order+1)
	postOrder(node.Right, lvl, order+1)
	if _, ok := lvl.m[order]; ok {
		l := len(lvl.m[order]) - 1
		lvl.m[order][l].Next = node
		lvl.m[order] = append(lvl.m[order], node)
	} else {
		lvl.m[order] = []*Node{node}
	}
}

func connect(root *Node) *Node {
	lvl := Level{map[int][]*Node{}}
	postOrder(root, &lvl, 0)
	return root
}
