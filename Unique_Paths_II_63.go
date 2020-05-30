package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	table := make([][]int, m)
	for i, _ := range table {
		table[i] = make([]int, n)
	}
	for i, _ := range table {
		if obstacleGrid[i][0] == 1 {
			break
		}
		table[i][0] = 1
	}
	for i, _ := range table[0] {
		if obstacleGrid[0][i] == 1 {
			break
		}
		table[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				table[i][j] = 0
			} else {
				table[i][j] = table[i-1][j] + table[i][j-1]
			}
		}
	}

	return table[m-1][n-1]
}
func main() {
	board := [][]int{{0, 0}}
	fmt.Println(uniquePathsWithObstacles(board))
}
