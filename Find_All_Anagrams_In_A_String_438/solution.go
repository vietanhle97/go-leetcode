package Find_All_Anagrams_In_A_String_438

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	m := len(s)
	n := len(p)
	count := make([]int, 26)
	for i := 0; i < n; i++ {
		count[p[i]-'a']++
	}
	res := make([]int, 0)

	for i := 0; i < m-n+1; i++ {
		cnt := make([]int, 26)
		copy(cnt, count)
		valid := true
		for j := 0; j < n; j++ {
			cnt[s[i+j]-'a']--
			if cnt[s[i+j]-'a'] < 0 {
				valid = false
				break
			}
		}
		if valid {
			res = append(res, i)
		}
	}
	return res
}
