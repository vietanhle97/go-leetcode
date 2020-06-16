package Balance_A_Binary_Search_Tree_1382

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrder(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left, nodes)
	*nodes = append(*nodes, root)
	inOrder(root.Right, nodes)
}

func buildTree(nodes []*TreeNode, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	node := nodes[mid]
	node.Left = buildTree(nodes, start, mid-1)
	node.Right = buildTree(nodes, mid+1, end)
	return node
}

func balanceBST(root *TreeNode) *TreeNode {
	nodes := make([]*TreeNode, 0)
	inOrder(root, &nodes)
	buildTree(nodes, 0, len(nodes)-1)
	if len(nodes)%2 == 0 {
		return nodes[len(nodes)/2-1]
	}
	return nodes[len(nodes)/2]

}
