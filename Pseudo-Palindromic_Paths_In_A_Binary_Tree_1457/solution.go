package Pseudo_Palindromic_Paths_In_A_Binary_Tree_1457

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func count(num []int) bool {
	cnt := 0
	for _, e := range num {
		if e%2 == 1 {
			cnt++
		}
		if cnt == 2 {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode, path []int, res *int) {
	if root == nil {
		return
	}
	path[root.Val-1]++
	if count(path) && root.Left == nil && root.Right == nil {
		*res++
	}
	ls, rs := make([]int, 9), make([]int, 9)
	copy(ls, path)
	copy(rs, path)
	dfs(root.Left, ls, res)
	dfs(root.Right, rs, res)
}

func pseudoPalindromicPaths(root *TreeNode) int {
	path, res := make([]int, 9), 0
	dfs(root, path, &res)
	return res

}
