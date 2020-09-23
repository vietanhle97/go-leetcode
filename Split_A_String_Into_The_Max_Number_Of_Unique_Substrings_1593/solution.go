package Split_A_String_Into_The_Max_Number_Of_Unique_Substrings_1593

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func exist(s string, path []string) bool {
	for _, e := range path {
		if s == e {
			return true
		}
	}
	return false
}

func dfs(start int, s string, path []string, res *int) {
	if start >= len(s) {
		*res = max(*res, len(path))
		return
	}
	for i := start; i < len(s); i++ {
		tmp := s[start : i+1]
		if !exist(tmp, path) {
			dfs(i+1, s, append(path, tmp), res)
		}
	}
}

func maxUniqueSplit(s string) int {
	res := 0
	dfs(0, s, []string{}, &res)
	return res
}
