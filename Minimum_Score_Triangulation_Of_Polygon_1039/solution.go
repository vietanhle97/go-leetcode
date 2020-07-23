package Minimum_Score_Triangulation_Of_Polygon_1039

const MaxInt = 1<<31 - 1

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minScoreTriangulation(A []int) int {
	m := len(A)

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for i := 2; i < m; i++ {
		for j := 0; j < m-i; j++ {
			st, en := j, i+j
			dp[st][en] = MaxInt
			for k := j + 1; k < en; k++ {
				dp[st][en] = min(dp[st][en], dp[st][k]+dp[k][en]+A[k]*A[st]*A[en])
			}
		}
	}
	return dp[0][m-1]
}
