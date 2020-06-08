package Palindrome_Partitioning_II_132

import (
	"math"
)

func isPalindrome(s string, m int) bool {
	i := 0
	j := m - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func minCut(s string) int {
	m := len(s)
	if s == "" {
		return 0
	}
	table := make([]float64, 0)
	for i := 0; i < m+1; i++ {
		table = append(table, math.Inf(1))
	}
	table[0] = 0
	for j := 1; j < m+1; j++ {
		k := 0
		for k < j {
			if isPalindrome(s[k:j], j-k) {
				table[j] = math.Min(table[k]+1, table[j])
			}
			k += 1
		}
	}
	return int(table[m] - 1)
}
