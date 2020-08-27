package Last_Stone_Weight_II_1049

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func lastStoneWeightII(stones []int) int {
	n, sum := len(stones), 0
	for i := range stones {
		sum += stones[i]
	}
	target := sum / 2
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < target+1; j++ {
			if stones[i-1] <= j {
				dp[i][j] = max(dp[i-1][j], stones[i-1]+dp[i-1][j-stones[i-1]])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return sum - dp[n][target]*2
}
