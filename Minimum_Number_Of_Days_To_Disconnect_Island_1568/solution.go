package Minimum_Number_Of_Days_To_Disconnect_Island_1568

var (
	m, n int
	vis  [][]bool
	poss [][]int
	X    = []int{-1, 1, 0, 0}
	Y    = []int{0, 0, -1, 1}
)

func valid(x, y int, grid [][]int) bool {
	return x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 && !vis[x][y]
}

func dfs(x, y int, grid [][]int, cnt *int) {
	vis[x][y] = true
	*cnt++
	for i := 0; i < 4; i++ {
		nx, ny := x+X[i], y+Y[i]
		if valid(nx, ny, grid) {
			dfs(nx, ny, grid, cnt)
		}
	}
}

func minDays(grid [][]int) int {
	m, n = len(grid), len(grid[0])
	poss, vis = make([][]int, 0), make([][]bool, m)
	for i := range grid {
		vis[i] = make([]bool, n)
		for j := range grid[i] {
			if grid[i][j] == 1 {
				poss = append(poss, []int{i, j})
			}
		}
	}
	if len(poss) < 2 {
		return 0
	}
	cnt := 0
	dfs(poss[0][0], poss[0][1], grid, &cnt)
	if cnt < len(poss) {
		return 0
	}

	for i := range poss {
		vis[poss[i][0]][poss[i][1]] = false
	}

	for i, p := range poss {
		cnt = 0
		for j := range poss {
			vis[poss[j][0]][poss[j][1]] = false
		}
		vis[p[0]][p[1]] = true
		if i == 0 {
			dfs(poss[i+1][0], poss[i+1][1], grid, &cnt)
			if cnt < len(poss)-1 {
				return 1
			}
			vis[p[0]][p[1]] = false
		} else {
			dfs(poss[i-1][0], poss[i-1][1], grid, &cnt)
			if cnt < len(poss)-1 {
				return 1
			}
			vis[p[0]][p[1]] = false
		}
	}
	return 2
}
