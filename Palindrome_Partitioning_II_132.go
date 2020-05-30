package main

import (
	"fmt"
	"math"
)

func isPalindrome(s string, m int) bool {
	i := 0
	j := m-1
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
	table := []float64{}
	for i:=0; i<m+1; i++ {
		table = append(table, math.Inf(1))
	}
	table[0] = 0
	for j:=1; j<m+1; j++ {
		for k:=0;k < j; k++ {
			if isPalindrome(s[k:j], j-k) {
				table[j] = math.Min(table[k] + 1, table[j])
			}
		}
	}
	return int(table[m] - 1)
}

func main() {
	s := "abbab"
	fmt.Println(minCut(s))
}
