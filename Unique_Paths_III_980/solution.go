package Unique_Paths_III_980

func check(p []int, path [][]int) bool {
	for _, e := range path {
		if e[0] == p[0] && e[1] == p[1] {
			return false
		}
	}
	return true
}

func isValid(x int, y int, m int, n int, grid [][]int) bool {
	if x < 0 || x >= m || y < 0 || y >= n {
		return false
	}
	if grid[x][y] == -1 {
		return false
	}
	return true
}

func DFS(row int, col int, m int, n int, grid [][]int, path [][]int, max_ int, cnt *int) {
	if grid[row][col] == 2 && len(path) == max_ {
		*cnt++
		return
	}
	if isValid(row-1, col, m, n, grid) && check([]int{row - 1, col}, path) {
		DFS(row-1, col, m, n, grid, append(path, []int{row, col}), max_, cnt)
	}
	if isValid(row+1, col, m, n, grid) && check([]int{row + 1, col}, path) {
		DFS(row+1, col, m, n, grid, append(path, []int{row, col}), max_, cnt)
	}
	if isValid(row, col-1, m, n, grid) && check([]int{row, col - 1}, path) {
		DFS(row, col-1, m, n, grid, append(path, []int{row, col}), max_, cnt)
	}
	if isValid(row, col+1, m, n, grid) && check([]int{row, col + 1}, path) {
		DFS(row, col+1, m, n, grid, append(path, []int{row, col}), max_, cnt)
	}
}

func uniquePathsIII(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	cnt := 0
	row, col := 0, 0
	obs := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				row, col = i, j
			}
			if grid[i][j] == -1 {
				obs += 1
			}
		}
	}
	DFS(row, col, m, n, grid, [][]int{}, m*n-obs-1, &cnt)
	return cnt
}
