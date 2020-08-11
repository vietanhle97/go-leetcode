package Maximal_Square_221

func MAX(a, b int) int {
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
func maximalSquareSol2(A [][]byte) int {
	m, ans := len(A), 0
	if m == 0 || len(A[0]) == 0 {
		return ans
	}
	n := len(A[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := range dp {
		for j := range dp[i] {
			if A[i][j] == '1' {
				dp[i][j] = 1
				if i > 0 && j > 0 {
					dp[i][j] += min(min(dp[i-1][j-1], dp[i][j-1]), dp[i-1][j])
				}
			}
			ans = max(ans, dp[i][j])
		}
	}
	return ans * ans
}
