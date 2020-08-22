package Count_Servers_That_Communicate_1267

func countServers(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	row, col, cnt := make([]int, m), make([]int, n), 0
	for i := 0; i < m; i++ {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				row[i]++
				col[j]++
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := range grid[i] {
			if grid[i][j] == 1 && (row[i] >= 2 || col[j] >= 2) {
				cnt++
			}
		}
	}
	return cnt
}
