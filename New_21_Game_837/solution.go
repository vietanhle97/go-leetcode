package New_21_Game_837

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func new21Game(N int, K int, W int) float64 {
	if K == 1 {
		if N >= W {
			return 1.0
		}
		return float64(N) / float64(W)
	}
	res := float64(min(N-K+1, W))
	dp := make([]float64, N+W+1)
	for k := K; k < N+1; k++ {
		dp[k] = 1.0
	}
	for k := K - 1; k > -1; k-- {
		dp[k] = res / float64(W)
		res += dp[k] - dp[k+W]
	}

	return dp[0]
}
