package Longest_Palindromic_Substring_5

func longestPalindrome(s string) string {
	m := len(s)
	max_ := 0
	start := 0
	end := 0
	if m < 2 {
		return s
	}
	table := make([][]bool, m)
	for i, _ := range table {
		table[i] = make([]bool, m)
		table[i][i] = true
	}
	for cl := 2; cl < m+1; cl++ {
		for i := 0; i < m-cl+1; i++ {
			j := i + cl - 1
			if s[i] == s[j] && cl == 2 {
				table[i][j] = true
			} else if s[i] == s[j] {
				table[i][j] = table[i+1][j-1]
			}
			if table[i][j] && j-i+1 > max_ {
				start = i
				end = j + 1
				max_ = j + 1 - i
			}

		}
	}
	if s[start:end] == "" {
		return string(s[0])
	}
	return s[start:end]
}
