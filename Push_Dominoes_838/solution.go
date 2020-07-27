package Push_Dominoes_838

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pushDominoes(s string) string {
	m := len(s)
	res := make([]int, m)
	cnt := 0
	ans := make([]rune, m)
	for i := range s {
		if s[i] == 'R' {
			cnt = m
		} else if s[i] == 'L' {
			cnt = 0
		} else {
			cnt = max(cnt-1, 0)
		}
		res[i] += cnt
	}
	cnt = 0
	for i := m - 1; i >= 0; i-- {
		if s[i] == 'L' {
			cnt = m
		} else if s[i] == 'R' {
			cnt = 0
		} else {
			cnt = max(cnt-1, 0)
		}
		res[i] -= cnt
	}
	for i, e := range res {
		if e == 0 {
			ans[i] = '.'
		} else if e > 0 {
			ans[i] = 'R'
		} else {
			ans[i] = 'L'
		}
	}
	// fmt.Println(res)
	return string(ans)
}
