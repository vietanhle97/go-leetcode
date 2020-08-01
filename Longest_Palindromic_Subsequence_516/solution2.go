package Longest_Palindromic_Subsequence_516

func MAX(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestPalindromeSubseqSol2(s string) int {
	m := len(s)

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
		dp[i][i] = 1
	}

	for cl := 2; cl < m+1; cl++ {
		for i := 0; i < m-cl+1; i++ {
			j := i + cl - 1
			if s[i] == s[j] && cl == 2 {
				dp[i][j] = 2
			} else if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = MAX(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][m-1]
}
