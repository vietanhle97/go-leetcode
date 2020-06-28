package Minimum_Falling_Path_Sum_931

import "math"

func isValid(x, y, m, n int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func minFallingPathSum(A [][]int) int {
	m := len(A)
	n := len(A[0])
	table := make([][]int, m)
	res := 100*100 + 1
	for i := 0; i < m; i++ {
		table[i] = make([]int, n)
		if i == 0 {
			for j := 0; j < n; j++ {
				table[i][j] = A[0][j]
				if i == m-1 {
					res = min(res, table[i][j])
				}
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			visited := false
			if isValid(i-1, j-1, m, n) {
				visited = true
				table[i][j] = table[i-1][j-1] + A[i][j]
			}
			if isValid(i-1, j, m, n) {
				if visited {
					table[i][j] = min(table[i][j], table[i-1][j]+A[i][j])
				} else {
					visited = true
					table[i][j] = table[i-1][j] + A[i][j]
				}
			}
			if isValid(i-1, j+1, m, n) {
				if visited {
					table[i][j] = min(table[i][j], table[i-1][j+1]+A[i][j])
				} else {
					visited = true
					table[i][j] = table[i-1][j+1] + A[i][j]
				}
			}
			if i == m-1 {
				res = min(res, table[i][j])
			}
		}
	}
	return res
}
