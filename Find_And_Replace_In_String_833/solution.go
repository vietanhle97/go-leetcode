package Find_And_Replace_In_String_833

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	m := map[int]string{}
	r := make([]int, len(S))
	res := ""
	for i := range r {
		r[i] = 1
	}
	for i, e := range indexes {
		m[e] = targets[i]
		tmp := e + len(sources[i])
		if sources[i] == S[e:tmp] {
			for j := e; j < tmp; j++ {
				r[j] = 0
			}
		}

	}

	for i := range r {
		if r[i] == 1 {
			res += string(S[i])
		} else {
			res += m[i]
		}
	}
	return res
}
