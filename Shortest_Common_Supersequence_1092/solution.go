package Shortest_Common_Supersequence_1092

func max(a ...int) int {
	res := a[0]
	for _, e := range a {
		if e > res {
			res = e
		}
	}
	return res
}
func shortestCommonSupersequence(s1 string, s2 string) string {

	m := len(s1)
	n := len(s2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
				continue
			}
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}

	}
	if dp[m][n] == 0 {
		return s1 + s2
	} else if dp[m][n] == m {
		return s2
	} else if dp[m][n] == n {
		return s1
	}
	ind, tmp := dp[m][n], dp[m][n]
	a, b := make([]int, ind), make([]int, ind)
	i, j := m, n
	for i > 0 && j > 0 {
		if s1[i-1] == s2[j-1] {
			a[tmp-1] = i - 1
			b[tmp-1] = j - 1
			i--
			j--
			tmp--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}
	res := ""
	i = 0
	if i == ind {
		res += string(s1[a[ind]])
		if a[ind]+1 < len(s1) {
			res += s1[a[ind]+1:]
		}
		if b[ind]+1 < len(s2) {
			res += s2[b[ind]+1:]
		}
		return res
	}
	for i < ind {
		if i == 0 {
			res += s1[:a[i]]
			res += s2[:b[i]]
			res += string(s1[a[i]])
		} else {
			res += s1[a[i-1]+1 : a[i]]
			res += s2[b[i-1]+1 : b[i]]
			res += string(s1[a[i]])
		}
		i++
	}

	if a[ind-1]+1 < len(s1) {
		res += s1[a[ind-1]+1:]
	}
	if b[ind-1]+1 < len(s2) {
		res += s2[b[ind-1]+1:]
	}

	return res
}
