package N_ary__Tree_Level_Order_Traversal_429

type Node struct {
	Val      int
	Children []*Node
}

type Mem struct {
	check map[int][]int
	order int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func preOrder(root *Node, order int, mem *Mem) {
	if root == nil {
		return
	}
	if _, ok := mem.check[order]; !ok {
		mem.check[order] = []int{}
	}
	mem.check[order] = append(mem.check[order], root.Val)
	mem.order = max(order, mem.order)
	for _, child := range root.Children {
		preOrder(child, order+1, mem)
	}

}

func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	mem := Mem{map[int][]int{}, 0}
	preOrder(root, 0, &mem)
	for i := 0; i < mem.order+1; i++ {
		res = append(res, mem.check[i])
	}
	return res
}
