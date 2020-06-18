package Longest_Substring_With_At_Least_K_Repeating_Characters_395

import (
	"math"
	"strings"
)

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func longestSubstring(s string, k int) int {
	alpha := make([]int, 26)
	for i := range s {
		alpha[s[i]-'a']++
	}

	diff := 0
	for i, v := range alpha {
		if v > 0 && v < k {
			diff = i + int('a')
			break
		}
	}
	if diff == 0 {
		return len(s)
	}
	S := strings.Split(s, string(diff))
	res := 0
	for _, v := range S {
		res = max(res, longestSubstring(v, k))
	}
	return res
}
