package Most_Stones_Removed_With_Same_Row_Or_Column_947

type node struct {
	i int
	j int
}

func dfs(a, b int, vis map[node]bool, m map[node][]node, cnt *int) {
	vis[node{a, b}] = true
	*cnt++
	for _, e := range m[node{a, b}] {
		if !vis[node{e.i, e.j}] {
			dfs(e.i, e.j, vis, m, cnt)
		}
	}
}

func removeStones(stones [][]int) int {
	m := map[node][]node{}
	for i, u := range stones {
		for j, v := range stones {
			if i != j && (u[0] == v[0] || u[1] == v[1]) {
				m[node{u[0], u[1]}] = append(m[node{u[0], u[1]}], node{v[0], v[1]})
				m[node{v[0], v[1]}] = append(m[node{v[0], v[1]}], node{u[0], u[1]})
			}
		}
	}
	vis := map[node]bool{}
	res := 0
	for _, e := range stones {
		cnt := 0
		if !vis[node{e[0], e[1]}] {
			dfs(e[0], e[1], vis, m, &cnt)
		}
		if cnt > 0 {
			res += cnt - 1
		}
	}
	return res
}
