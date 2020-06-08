package Maximum_Width_Of_Binary_Tree_662

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Mem struct {
	check map[int][]int
	max   int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func preOrder(root *TreeNode, order, cur int, mem *Mem) {
	if root == nil {
		return
	}
	if _, ok := mem.check[order]; !ok {
		mem.check[order] = []int{cur, cur}
	} else {
		mem.check[order][1] = max(mem.check[order][1], cur)
	}
	mem.max = max(mem.check[order][1]-mem.check[order][0]+1, mem.max)
	preOrder(root.Left, order+1, 2*cur-1, mem)
	preOrder(root.Right, order+1, 2*cur, mem)
}

func widthOfBinaryTree(root *TreeNode) int {
	mem := Mem{map[int][]int{}, 0}
	preOrder(root, 0, 1, &mem)
	return mem.max
}
