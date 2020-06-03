package Path_Sum_II_113

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, path []int, target int, res *[][]int) {
	if root.Left == nil && root.Right == nil {
		if target == root.Val {
			path = append(path, root.Val)
			*res = append(*res, append([]int{}, path...))
		}
		return
	}
	if root.Left != nil {
		dfs(root.Left, append(path, root.Val), target-root.Val, res)
	}
	if root.Right != nil {
		dfs(root.Right, append(path, root.Val), target-root.Val, res)
	}

}

func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	dfs(root, []int{}, sum, &res)
	return res
}
