package Binary_Tree_Coloring_Game_1145

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNode(root *TreeNode, cur []int, table *map[int]int, x int, X *TreeNode) {
	if root == nil {
		return
	}
	if root.Val == x {
		X.Val = x
		X.Left = root.Left
		X.Right = root.Right
	}
	left, right := []int{0}, []int{0}
	countNode(root.Left, left, table, x, X)
	countNode(root.Right, right, table, x, X)
	cur[0] = left[0] + right[0] + 1
	(*table)[root.Val] = cur[0]
}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	cur := []int{0}
	table := map[int]int{}
	X := TreeNode{}
	countNode(root, cur, &table, x, &X)
	if x == root.Val {
		left, right := 0, 0
		if root.Left != nil {
			left = table[root.Left.Val]
		}
		if root.Right != nil {
			right = table[root.Right.Val]
		}
		return 1+left < right || 1+right < left
	} else {
		left, right := 0, 0
		if X.Left != nil {
			left = table[X.Left.Val]
		}
		if X.Right != nil {
			right = table[X.Right.Val]
		}
		return table[root.Val]-table[x] > table[x] || left > table[root.Val]-left || right > table[root.Val]-right
	}
}
