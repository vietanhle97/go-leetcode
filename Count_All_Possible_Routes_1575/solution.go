package Count_All_Possible_Routes_1575

const mod = int(1e9) + 7

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dfs(u, fuel, end int, g []int, dp [][]int) int {
	if dp[u][fuel] != -1 {
		return dp[u][fuel]
	}
	res := 0
	if u == end {
		res = 1
	}
	for v, e := range g {
		if v != u && abs(g[u]-g[v]) <= fuel {
			res += dfs(v, fuel-abs(e-g[u]), end, g, dp)
			res %= mod
		}
	}
	dp[u][fuel] = res
	return res
}

func countRoutes(g []int, start int, finish int, fuel int) int {
	m := len(g)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, fuel+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return dfs(start, fuel, finish, g, dp) % mod

}
