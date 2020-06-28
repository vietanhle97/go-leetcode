package Permutation_In_String_567

func checkInclusion(p string, s string) bool {
	if len(p) > len(s) {
		return false
	}
	m := len(s)
	n := len(p)
	count := make([]int, 26)
	for i := 0; i < n; i++ {
		count[p[i]-'a']++
	}

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
			return valid
		}
	}
	return false
}
