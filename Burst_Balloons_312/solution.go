package Burst_Balloons_312

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxCoins(A []int) int {
	m := len(A)
	if m == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for L := m - 1; L >= 0; L-- {
		for R := L; R < m; R++ {
			for i := L; i <= R; i++ {
				a, b, c, d := 0, 0, 0, 0
				if L == 0 {
					a = 1
				} else {
					a = A[L-1]
				}
				if R == m-1 {
					b = 1
				} else {
					b = A[R+1]
				}
				if L > i-1 {
					c = 0
				} else {
					c = dp[L][i-1]
				}
				if R < i+1 {
					d = 0
				} else {
					d = dp[i+1][R]
				}
				dp[L][R] = max(dp[L][R], A[i]*a*b+c+d)
			}
		}
	}
	return dp[0][m-1]
}
