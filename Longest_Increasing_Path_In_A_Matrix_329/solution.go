package Longest_Increasing_Path_In_A_Matrix_329

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfsInc(i, j int, g [][]int, vis [][]int) int {
	if vis[i][j] != 0 {
		return vis[i][j]
	}
	res := 1
	if i > 0 && g[i-1][j] > g[i][j] {
		res = max(res, 1+dfsInc(i-1, j, g, vis))
	}
	if j > 0 && g[i][j-1] > g[i][j] {
		res = max(res, 1+dfsInc(i, j-1, g, vis))
	}
	if i < len(g)-1 && g[i+1][j] > g[i][j] {
		res = max(res, 1+dfsInc(i+1, j, g, vis))
	}
	if j < len(g[0])-1 && g[i][j+1] > g[i][j] {
		res = max(res, 1+dfsInc(i, j+1, g, vis))
	}
	vis[i][j] = res
	return res
}

func longestIncreasingPath(matrix [][]int) int {
	res := 0
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	vis := make([][]int, m)
	for i := range vis {
		vis[i] = make([]int, n)
	}
	for i := range matrix {
		for j := range matrix[0] {
			res = max(res, dfsInc(i, j, matrix, vis))
		}
	}
	return res
}
