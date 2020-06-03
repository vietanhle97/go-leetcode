package Word_Break_139

func wordBreak(s string, wordDict []string) bool {
	m := len(s)
	table := make([][]bool, m)
	check := map[string]bool{}
	for _, e := range wordDict {
		check[e] = true
	}
	for i, _ := range table {
		table[i] = make([]bool, m)
	}
	for i := 0; i < m; i++ {
		for j := i; j < m; j++ {
			if i == 0 {
				tmp := s[i : j+1]
				if _, ok := check[tmp]; ok {
					if j == m-1 {
						return true
					} else {
						table[i][j] = true
					}
				}
			} else {
				for k := 0; k < i; k++ {
					if table[k][i-1] {
						tmp := s[i : j+1]
						if _, ok := check[tmp]; ok {
							if j == m-1 {
								return true
							} else {
								table[i][j] = true
							}
						}
					}
				}
			}
		}
	}
	return false
}
