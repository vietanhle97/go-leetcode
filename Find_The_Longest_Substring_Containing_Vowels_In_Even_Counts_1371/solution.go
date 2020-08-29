package Find_The_Longest_Substring_Containing_Vowels_In_Even_Counts_1371

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func findTheLongestSubstring(s string) int {
	m := map[byte]int{'a': 0, 'e': 1, 'i': 2, 'o': 3, 'u': 4}
	cnt := make([]int, 32)
	for i := range cnt {
		cnt[i] = len(s)
	}
	cnt[0] = -1
	mask, res := 0, 0
	for i := range s {
		if _, ok := m[s[i]]; ok {
			mask ^= 1 << (m[s[i]])
		}
		for j := 0; j < 32; j++ {
			if mask^j == 0 {
				res = max(res, i-cnt[mask])
			}
		}
		cnt[mask] = min(cnt[mask], i)
	}
	return res
}
