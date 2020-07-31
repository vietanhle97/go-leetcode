package Find_Largest_Value_In_Each_Tree_Row_515

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(root *TreeNode, lv int, m map[int][]int, max *int) {
	if root == nil {
		return
	}
	if _, ok := m[lv]; !ok {
		m[lv] = append(m[lv], root.Val)
	} else {
		if root.Val > m[lv][len(m[lv])-1] {
			m[lv] = append(m[lv], root.Val)
		}
	}
	if lv > *max {
		*max = lv
	}
	preOrder(root.Left, lv+1, m, max)
	preOrder(root.Right, lv+1, m, max)
}
func largestValues(root *TreeNode) []int {
	max, m, res := 0, map[int][]int{}, make([]int, 0)
	if root == nil {
		return res
	}
	preOrder(root, 0, m, &max)
	for i := 0; i <= max; i++ {
		res = append(res, m[i][len(m[i])-1])
	}
	return res
}
