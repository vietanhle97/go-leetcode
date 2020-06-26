package Minimum_Path_Sum_64

import "math"

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0

	}
	table := make([][]int, m)
	for i := range table {
		table[i] = make([]int, n)
		if i == 0 {
			table[0][0] = grid[0][0]
			for j := 1; j < n; j++ {
				table[i][j] = table[i][j-1] + grid[i][j]
			}
		} else {
			table[i][0] = table[i-1][0] + grid[i][0]
			for j := 1; j < n; j++ {
				table[i][j] = min(table[i-1][j], table[i][j-1]) + grid[i][j]
			}
		}
	}
	return table[m-1][n-1]
}
