package Maximum_Score_From_Removing_Substrings_1717

func maximumGain(s string, x int, y int) int {
	st, st2 := make([]byte, 0), make([]byte, 0)
	res := 0
	cmp, v1, v2 := []byte{'a', 'b'}, x, y
	if y > x {
		cmp, v1, v2 = []byte{'b', 'a'}, y, x
	}
	for i := range s {
		if len(st) == 0 {
			st = append(st, s[i])
		} else if st[len(st)-1] == cmp[0] && s[i] == cmp[1] {
			res += v1
			st = st[:len(st)-1]
		} else {
			st = append(st, s[i])
		}
	}

	for i := range st {
		if len(st2) == 0 {
			st2 = append(st2, st[i])
		} else if st2[len(st2)-1] == cmp[1] && st[i] == cmp[0] {
			res += v2
			st2 = st2[:len(st2)-1]
		} else {
			st2 = append(st2, st[i])
		}
	}

	return res
}
