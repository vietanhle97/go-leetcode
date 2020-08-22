package Coin_Change_322

const MaxInt = int(1e8 + 5)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = MaxInt
	}
	dp[0] = 0

	for _, coin := range coins {
		for x := coin; x < amount+1; x++ {
			dp[x] = min(dp[x], dp[x-coin]+1)
		}

	}
	if dp[amount] == MaxInt {
		return 0
	}
	return dp[amount]
}
