package Word_Break_139

func wordBreak(s string, wordDict []string) bool {
	m := len(s)
	table := make([]bool, m)
	check := map[string]bool{}
	for _, e := range wordDict {
		check[e] = true
	}

	for i := range table {
		if check[s[0:i+1]] {
			if i == m-1 {
				return true
			}
			table[i] = true
		}
	}
	for i := 1; i < m; i++ {
		if table[i] {
			continue
		}
		for k := 0; k < i; k++ {
			if table[k] && check[s[k+1:i+1]] {
				if i == m-1 {
					return true
				}
				table[i] = true
			}
		}
	}
	return false
}
