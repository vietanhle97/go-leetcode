package Stone_Game_V_1563

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func dfs(i, j int, dp [][]int, P []int) int {
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	res := 0
	for k := i; k < j; k++ {
		left, right := P[k+1]-P[i], P[j+1]-P[k+1]
		if left <= right {
			res = max(res, dfs(i, k, dp, P)+left)
		}
		if right <= left {
			res = max(res, dfs(k+1, j, dp, P)+right)
		}
	}
	dp[i][j] = res
	return dp[i][j]
}

func stoneGameV(S []int) int {
	P := []int{0}
	for i := range S {
		P = append(P, P[i]+S[i])
	}
	dp := make([][]int, len(S))
	for i := range dp {
		dp[i] = make([]int, len(S))
	}
	return dfs(0, len(S)-1, dp, P)
}
