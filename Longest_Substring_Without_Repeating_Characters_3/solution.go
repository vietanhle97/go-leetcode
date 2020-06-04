package Longest_Substring_Without_Repeating_Characters_3

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	check := map[uint8]int{}
	start := 0
	end := 0
	res := 0
	for i := range s {
		if _, ok := check[s[i]]; !ok {
			check[s[i]] = i
			end = i
			res = max(res, end-start+1)
		} else {
			start = max(start, check[s[i]]+1)
			end = i
			check[s[i]] = i
			res = max(res, end-start+1)
		}
	}
	return res
}
