package Maximum_Level_Sum_Of_A_Binary_Tree_1161

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func preOrder(root *TreeNode, cur int, level map[int]int) {
	if root == nil {
		return
	}
	level[cur] += root.Val
	preOrder(root.Left, cur+1, level)
	preOrder(root.Right, cur+1, level)
}

func maxLevelSum(root *TreeNode) int {
	res, MAX := 0, 0
	level := map[int]int{}
	preOrder(root, 1, level)
	for k, v := range level {
		if v > MAX {
			res, MAX = k, v
		} else if v == MAX {
			res = min(res, k)
		}
	}
	return res

}
