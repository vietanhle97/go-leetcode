package Delete_Nodes_And_Return_Forest_1110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(root, parent *TreeNode, left bool, del map[int]bool, res *[]*TreeNode) {
	if root == nil {
		return
	}
	// fmt.Println(root.Val)
	if del[root.Val] {
		if root.Left != nil {
			if _, ok := del[root.Left.Val]; !ok {
				*res = append(*res, root.Left)
			}
		}
		if root.Right != nil {
			if _, ok := del[root.Right.Val]; !ok {
				*res = append(*res, root.Right)
			}
		}
		if parent != nil {
			if left {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		}
	}

	preOrder(root.Left, root, true, del, res)
	preOrder(root.Right, root, false, del, res)
}

func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	m := map[int]bool{}
	for _, e := range toDelete {
		m[e] = true
	}
	// fmt.Println(m)
	res := make([]*TreeNode, 0)
	preOrder(root, nil, false, m, &res)
	if _, ok := m[root.Val]; !ok {
		res = append(res, root)
	}
	return res

}
