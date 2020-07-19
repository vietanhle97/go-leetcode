package Stone_Game_III_1406

func MAX(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func stoneGameIIISol2(piles []int) string {
	N := len(piles)
	sum := make([]int, N+1)
	sum[N] = 0
	for i := N - 1; i >= 0; i-- {
		sum[i] = sum[i+1] + piles[i]
	}
	dp := make([]int, N+1)
	dp[N] = 0
	for i := N - 1; i >= 0; i-- {
		dp[i] = piles[i] + sum[i+1] - dp[i+1]
		for j := i; j < i+3 && j < N; j++ {
			dp[i] = MAX(dp[i], sum[i]-dp[j+1])
		}
	}
	// fmt.Println(dp,sum)

	if dp[0]*2 > sum[0] {
		return "Alice"
	} else if dp[0]*2 == sum[0] {
		return "Tie"
	}
	return "Bob"
}
