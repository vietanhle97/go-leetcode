package Path_With_Minimum_Effort_1631

var x = []int{-1, 1, 0, 0}
var y = []int{0, 0, -1, 1}
var m, n int

func valid(x, y int) bool {
	return x < y && x >= 0
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
func dfs(i, j, k int, g [][]int, vis [][]int) bool {
	vis[i][j] = 1
	if i == m-1 && j == n-1 {
		return true
	}
	for d := 0; d < 4; d++ {
		nx, ny := i+x[d], j+y[d]
		if valid(nx, m) && valid(ny, n) && vis[nx][ny] == 0 && abs(g[nx][ny]-g[i][j]) <= k {
			if dfs(nx, ny, k, g, vis) {
				return true
			}
		}
	}
	return false
}

func minimumEffortPath(g [][]int) int {
	m = len(g)
	n = len(g[0])
	res := 1<<31 - 1
	vis := make([][]int, m)
	l, r := 0, int(1e6)
	for l <= r {
		for i := range vis {
			vis[i] = make([]int, n)
		}
		mid := l + (r-l)/2
		tmp := dfs(0, 0, mid, g, vis)
		if tmp {
			res = min(res, mid)
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}
