package Minimum_Cost_To_Merge_Stones_1000

const MaxInt = 1<<31 - 1

var L, R int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mergeStones(S []int, K int) int {
	m := len(S)
	if m == 1 {
		return 0
	}
	if (m-K)%(K-1) != 0 || K > m {
		return -1
	}
	dp, P := make([][]int, m+1), make([]int, m+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i < m; i++ {
		P[i+1] = P[i] + S[i]
	}

	for i := K - 1; i < m; i++ {
		L = 0
		for L, R = 0, L+i; L < m && R < m; L, R = L+1, R+1 {
			if i == K-1 {
				dp[L][R] = P[R+1] - P[L]
				continue
			}
			dp[L][R] = MaxInt
			for j := L; j < R; j += K - 1 {
				dp[L][R] = min(dp[L][R], dp[L][j]+dp[j+1][R])
			}
			if i%(K-1) == 0 {
				dp[L][R] += P[R+1] - P[L]
			}
		}
	}
	return dp[0][m-1]
}
