package Path_Sum_III_437

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(root *TreeNode, target int, path []int, res *int) {
	if root == nil {
		return
	}
	if root.Val == target {
		*res++
	}
	sum := root.Val
	for i := len(path) - 1; i > -1; i-- {
		sum += path[i]
		if sum == target {
			*res++
		}
	}
	preOrder(root.Left, target, append(path, root.Val), res)
	preOrder(root.Right, target, append(path, root.Val), res)
}

func pathSum(root *TreeNode, sum int) int {
	res := 0
	preOrder(root, sum, []int{}, &res)
	return res
}
