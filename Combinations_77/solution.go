package Combinations_77

func dfs(start int, n int, k int, res *[][]int, path []int) {
	if len(path) == k {
		*res = append(*res, append([]int{}, path...))
	}
	for i := start; i < n+1; i++ {
		dfs(i+1, n, k, res, append(path, i))
	}
}

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	dfs(1, n, k, &res, []int{})
	return res
}
