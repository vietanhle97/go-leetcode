package Stone_Game_877

// This is dp solution since I want to practice dp.
// Clearly we can see that the first one can always win
// so the optimal solution is as follow:

/* func stoneGame(piles []int) bool {
	return true
}
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func stoneGame(piles []int) bool {

	N := len(piles)
	dp := make([][]int, N+2)
	for i := range dp {
		dp[i] = make([]int, N+2)
	}
	for i := 1; i < N+1; i++ {
		for j, k := 0, i-1; k < N; j, k = j+1, k+1 {
			p := (j + k + N) % 2
			if p == 1 {
				dp[j+1][k+1] = max(piles[j]+dp[j+2][k+1], piles[k]+dp[j+1][k])
			} else {
				dp[j+1][k+1] = min(-piles[j]+dp[j+2][k+1], -piles[k]+dp[j+1][k])
			}
		}
	}
	// fmt.Println(dp)
	return dp[1][N] > 0
}
