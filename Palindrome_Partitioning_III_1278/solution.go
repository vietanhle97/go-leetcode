package Palindrome_Partitioning_III_1278

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func need(s string, i, j int) int {
	res := 0
	for i < j {
		if s[i] != s[j] {
			res++
		}
		i++
		j--
	}
	return res
}

func dfs(c [][]int, s string, i, k int) int {
	n := len(s)

	if i >= n || k == 0 {
		return 0
	}

	if k == 1 {
		return need(s, i, n-1)
	}

	if c[i][k] >= 0 {
		return c[i][k]
	}

	res := n
	for j := i; j < n-k+1; j++ {
		res = min(res, need(s, i, j)+dfs(c, s, j+1, k-1))
	}
	c[i][k] = res
	return res
}

func palindromePartition(s string, k int) int {
	n := len(s)
	c := make([][]int, n)
	for i := range c {
		c[i] = make([]int, n+1)
		for j := range c[i] {
			c[i][j] = -1
		}
	}
	return dfs(c, s, 0, k)
}
