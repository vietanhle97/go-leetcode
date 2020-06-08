package Most_Frequent_Subtree_Sum_508

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Mem struct {
	check map[int]int
	MAX   int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func postOrder(node *TreeNode, cur []int, mem *Mem) {
	if node == nil {
		return
	}
	left, right := []int{0}, []int{0}
	postOrder(node.Left, left, mem)
	postOrder(node.Right, right, mem)
	cur[0] = left[0] + right[0] + node.Val
	if _, ok := mem.check[cur[0]]; !ok {
		mem.check[cur[0]] = 1
	} else {
		mem.check[cur[0]] += 1
	}
	mem.MAX = max(mem.check[cur[0]], mem.MAX)
}

func findFrequentTreeSum(root *TreeNode) []int {
	res := make([]int, 0)
	mem := Mem{map[int]int{}, 0}
	cur := []int{0}
	postOrder(root, cur, &mem)
	for k, v := range mem.check {
		if v == mem.MAX {
			res = append(res, k)
		}
	}
	return res
}
