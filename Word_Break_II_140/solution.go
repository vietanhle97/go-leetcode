package Word_Break_II_140

func canPartition(s string, wordDict []string) (bool, []bool, map[string]bool) {
	m := len(s)
	table := make([]bool, m)
	check := map[string]bool{}
	for _, e := range wordDict {
		check[e] = true
	}

	for i := range table {
		if check[s[0:i+1]] {
			table[i] = true
		}
	}
	for i := 1; i < m; i++ {
		if table[i] {
			continue
		}
		for k := 0; k < i; k++ {
			if table[k] && check[s[k+1:i+1]] {
				table[i] = true
			}
		}
	}
	return table[m-1], table, check
}

func wordBreak(s string, wordDict []string) []string {
	b, table, check := canPartition(s, wordDict)
	if !b {
		return []string{}
	}
	m := len(s)
	res := make([]string, 0)
	words := map[int][]string{}

	for i := range table {
		tmp := s[0 : i+1]
		if check[tmp] {
			if i == m-1 {
				res = append(res, s)
			} else {
				if _, ok := words[i]; !ok {
					words[i] = []string{tmp + " "}
				}
			}

		}
	}
	for i := 1; i < m; i++ {
		for k := 0; k < i; k++ {
			tmp := s[k+1 : i+1]
			if table[k] && check[tmp] {
				if i == m-1 {
					for i := range words[k] {
						res = append(res, words[k][i]+tmp)
					}
				} else {
					if _, ok := words[i]; !ok {
						words[i] = []string{}
					}
					for j := range words[k] {
						words[i] = append(words[i], words[k][j]+tmp+" ")
					}
				}
			}
		}
	}
	return res
}
