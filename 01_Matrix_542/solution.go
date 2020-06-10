package _1_Matrix_542

import "math"

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func updateMatrix(matrix [][]int) [][]int {
	MAX := 50000
	m := len(matrix)
	if m == 0 {
		return matrix
	}
	n := len(matrix[0])
	if n == 0 {
		return matrix
	}
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				continue
			} else {
				res[i][j] = MAX
				if i > 0 {
					res[i][j] = min(res[i][j], res[i-1][j]+1)
				}
				if j > 0 {
					res[i][j] = min(res[i][j], res[i][j-1]+1)
				}
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i < m-1 {
				res[i][j] = min(res[i][j], res[i+1][j]+1)
			}
			if j < n-1 {
				res[i][j] = min(res[i][j], res[i][j+1]+1)
			}
		}
	}
	return res
}
