package Count_Sorted_Vowel_Strings_1641

func countVowelStrings(n int) int {
	dp := make([][]int, n)
	dp[0] = []int{1, 1, 1, 1, 1}
	for i := 1; i < n; i++ {
		dp[i] = make([]int, 5)
		for j := 4; j >= 0; j-- {
			if j == 4 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j+1] + dp[i-1][j]
			}

		}
	}
	res := 0
	for i := range dp[n-1] {
		res += dp[n-1][i]
	}
	return res
}
