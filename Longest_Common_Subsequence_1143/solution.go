package Longest_Common_Subsequence_1143

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= c {
		return b
	}
	return c
}
func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

	if m == 0 || n == 0 {
		return 0
	}

	table := make([][]int, m+1)
	for i := range table {
		table[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i-1] == text2[j-1] {
				table[i][j] = table[i-1][j-1] + 1
			} else {
				table[i][j] = max(table[i-1][j], table[i-1][j-1], table[i][j-1])
			}
		}
	}
	return table[m][n]
}
