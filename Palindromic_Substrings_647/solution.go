package Palindromic_Substrings_647

func findPalindrome(s string) [][]bool {
	m := len(s)
	table := make([][]bool, m)
	for i := 0; i < m; i++ {
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
			} else {
				table[i][j] = false
			}
		}
	}
	return table
}

func countSubstrings(s string) int {
	m := len(s)
	res := 0
	if m == 1 {
		return 1
	}
	table := make([][]int, m)
	palin := findPalindrome(s)
	for i := 0; i < m; i++ {
		table[i] = make([]int, m)
		table[i][i] = 1
	}
	for i := 0; i < m; i++ {
		for j := i; j < m; j++ {
			if j == i {
				continue
			}
			if !palin[i][j] {
				table[i][j] = table[i][j-1]
			} else {
				table[i][j] = table[i][j-1] + 1
			}
		}
	}
	for i := 0; i < m; i++ {
		res += table[i][m-1]
	}
	return res
}
