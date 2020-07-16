package Knight_Dialer_935

func knightDialer(N int) int {
	MOD := int(1e9 + 7)
	moves := [][]int{{4, 6}, {6, 8}, {7, 9}, {4, 8}, {3, 9, 0}, {}, {1, 7, 0}, {2, 6}, {1, 3}, {2, 4}}

	dp := [10]int{}
	for i := range dp {
		dp[i] = 1
	}
	for hops := 0; hops < N-1; hops++ {
		dp2 := [10]int{}
		for node, count := range dp {
			for _, nei := range moves[node] {
				dp2[nei] += count
				dp2[nei] %= MOD
			}
		}
		dp = dp2
	}
	res := 0
	for _, e := range dp {
		res += e
	}
	return res % MOD
}
