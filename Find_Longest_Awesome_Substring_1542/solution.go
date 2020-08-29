package Find_Longest_Awesome_Substring_1542

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

func longestAwesome(s string) int {
	m := make([]int, 1024)
	for i := range m {
		m[i] = len(s)
	}
	m[0] = -1
	mask, res := 0, 0
	for i := range s {
		mask ^= 1 << (s[i] - '0')
		res = max(res, i-m[mask])
		for j := 0; j < 10; j++ {
			test := mask ^ (1 << j)
			res = max(res, i-m[test])
		}
		m[mask] = min(m[mask], i)
	}
	// fmt.Println(m)
	return res
}
