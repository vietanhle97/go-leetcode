package Letter_Tile_Possibilities_1079

var m map[string]bool

func backtrack(s, path string, res *int) {
	if _, ok := m[path]; !ok && len(path) > 0 {
		*res++
		m[path] = true
	}
	for i := range s {
		backtrack(s[:i]+s[i+1:], path+string(s[i]), res)
	}
}

func numTilePossibilities(s string) int {
	m = map[string]bool{}
	res := 0
	backtrack(s, "", &res)
	return res
}
