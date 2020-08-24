package Cherry_Pickup_741

const MinInt = -(1<<31 - 1) - 1

var (
	g  [][]int
	dp [][][]int
	N  int
)

func max(a ...int) int {
	res := a[0]
	for _, e := range a {
		if e > res {
			res = e
		}
	}
	return res
}

func dfs(r1, c1, c2 int) int {
	r2 := r1 + c1 - c2
	ans := MinInt
	if r1 == N || c1 == N || c2 == N || r2 == N || g[r1][c1] == -1 || g[r2][c2] == -1 {
		return MinInt
	} else if r1 == c1 && c1 == N-1 {
		return g[r1][c1]
	} else if dp[r1][c1][c2] != MinInt {
		return dp[r1][c1][c2]
	} else {
		ans = g[r1][c1] + g[r2][c2]
		if c1 == c2 {
			ans -= g[r2][c2]
		}
		ans += max(dfs(r1+1, c1, c2), dfs(r1, c1+1, c2), dfs(r1+1, c1, c2+1), dfs(r1, c1+1, c2+1))
	}
	dp[r1][c1][c2] = ans
	return ans
}

func cherryPickup(grid [][]int) int {
	N = len(grid)
	g = grid
	dp = make([][][]int, N)
	for i := range dp {
		dp[i] = make([][]int, N)
		for j := range dp[i] {
			dp[i][j] = make([]int, N)
			for k := range dp[i][j] {
				dp[i][j][k] = MinInt
			}
		}
	}
	return max(0, dfs(0, 0, 0))
}
