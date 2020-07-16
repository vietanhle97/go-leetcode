package Interleaving_String_97

func isInterleave(s1 string, s2 string, s3 string) bool {

	m, n := len(s1), len(s2)
	if len(s3) != m+n {
		return false
	}
	dp := make([][]bool, m+n+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, m+n+1)
		for j := 0; j <= n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1]
			} else {
				dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}
	return dp[m][n]

}
