package Remove_Duplicate_Letters_316

func exist(ch byte, st []byte) bool {
	for _, c := range st {
		if ch == c {
			return true
		}
	}
	return false
}

func removeDuplicateLetters(s string) string {
	last, st, res := make([]int, 26), make([]byte, 0), ""
	for i := range s {
		last[s[i]-'a'] = i
	}

	for i := range s {
		if exist(s[i], st) {
			continue
		}
		for len(st) > 0 && st[len(st)-1] > s[i] && i < last[st[len(st)-1]-'a'] {
			st = st[:len(st)-1]
		}
		st = append(st, s[i])
	}
	for _, e := range st {
		res += string(e)
	}
	return res
}
