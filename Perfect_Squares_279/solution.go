package Perfect_Squares_279

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func numSquares(n int) int {
	m := []int{0}
	for i := 1; i*i <= n; i++ {
		m = append(m, i*i)

	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for k := 1; k < len(m) && m[k] <= i; k++ {
			if i-m[k] > 0 {
				if dp[i] == 0 {
					dp[i] = dp[i-m[k]] + 1
				}
				dp[i] = min(dp[i], dp[i-m[k]]+1)
			} else {
				dp[i] = 1
			}
		}
	}
	return dp[n]
}
