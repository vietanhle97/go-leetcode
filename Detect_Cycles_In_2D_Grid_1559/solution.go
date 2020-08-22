package Detect_Cycles_In_2D_Grid_1559

var v [][]int
var X = []int{-1, 1, 0, 0}
var Y = []int{0, 0, -1, 1}
var m, n int
var res = false

func dfs(x, y, pa, pb int, c byte, g [][]byte) bool {
	v[x][y] = int(c)
	for i := 0; i < 4; i++ {
		a, b := x+X[i], y+Y[i]
		if a < m && a >= 0 && b < n && b >= 0 {
			if v[a][b] == 0 && g[a][b] == c {
				if dfs(a, b, x, y, c, g) {
					return true
				}
			} else if v[a][b] == int(c) && a != pa && b != pb {
				return true
			}
		}
	}
	return false
}

func containsCycle(grid [][]byte) bool {
	m, n = len(grid), len(grid[0])
	v = make([][]int, m)
	for i := range v {
		v[i] = make([]int, n)
		for j := range v[i] {
			v[i][j] = 0
		}
	}
	for i := range grid {
		for j := range grid[i] {
			if v[i][j] == 0 {
				if dfs(i, j, -1, -1, grid[i][j], grid) {
					return true
				}
			}
		}
	}
	return false
}
