package Bulb_Switcher_IV_1529

func minFlips(s string) int {
	m := len(s)
	st := -1
	for i := range s {
		if s[i] == '1' {
			st = i
			break
		}
	}
	if st == -1 {
		return 0
	}
	cur := s[st]
	cnt := 1
	for i := st + 1; i < m; i++ {
		if s[i] != cur {
			cnt++
			cur = s[i]
		}
	}
	return cnt
}
