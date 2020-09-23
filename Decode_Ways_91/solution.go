package Decode_Ways_91

import "strconv"

func numDecodings(s string) int {
	m := map[string]int{}
	for i := 1; i <= 26; i++ {
		m[strconv.Itoa(i)] = i
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 0; i < len(s); i++ {
		if i == 0 {
			if _, ok := m[string(s[i])]; ok {
				dp[i+1]++
			}
		} else {
			if _, ok := m[s[i-1:i+1]]; ok {
				dp[i+1] += dp[i-1]
			}
			if _, ok := m[string(s[i])]; ok {
				dp[i+1] += dp[i]
			}
		}
	}
	return dp[len(s)]
}
