package All_Elements_In_Two_Binary_Search_Trees_1305

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, res)
	*res = append(*res, root.Val)
	dfs(root.Right, res)
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	b1, b2 := make([]int, 0), make([]int, 0)
	dfs(root1, &b1)
	dfs(root2, &b2)
	i, j, res := 0, 0, make([]int, 0)
	for i < len(b1) || j < len(b2) {
		if i < len(b1) && j < len(b2) && b1[i] <= b2[j] {
			res = append(res, b1[i])
			i++
		} else if j < len(b2) && i < len(b1) && b2[j] < b1[i] {
			res = append(res, b2[j])
			j++
		} else {
			if i < len(b1) {
				res = append(res, b1[i])
				i++
			}
			if j < len(b2) {
				res = append(res, b2[j])
				j++
			}
		}
	}
	return res
}
