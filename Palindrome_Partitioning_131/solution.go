package Palindrome_Partitioning_131

func backtrack(start int, s string, res *[][]string, path []string, palin [][]bool) {
	if start >= len(s) {
		*res = append(*res, append([]string{}, path...))
		return
	}

	for i := start; i < len(s); i++ {
		if palin[start][i] {
			backtrack(i+1, s, res, append(path, s[start:i+1]), palin)
		}
	}
}

func partition(s string) [][]string {
	m := len(s)
	if m == 1 {
		return [][]string{{s}}
	}
	res := make([][]string, 0)
	palin := make([][]bool, m)
	for i := 0; i < m; i++ {
		palin[i] = make([]bool, m)
		palin[i][i] = true
	}

	for cl := 2; cl < m+1; cl++ {
		for i := 0; i < m-cl+1; i++ {
			j := i + cl - 1
			if s[i] == s[j] && cl == 2 {
				palin[i][j] = true
			} else if s[i] == s[j] {
				palin[i][j] = palin[i+1][j-1]
			} else {
				palin[i][j] = false
			}
		}
	}

	backtrack(0, s, &res, []string{}, palin)
	return res
}
