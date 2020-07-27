package Course_Schedule_IV_1462

func dfs(root, v int, vis []bool, g [][]int, sp [][]bool) {
	vis[v] = true
	for _, e := range g[v] {
		sp[root][e] = true
		if !vis[e] {
			dfs(root, e, vis, g, sp)
		}
	}
}
func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
	g := make([][]int, n)
	sp := make([][]bool, n)
	res := make([]bool, 0)
	for _, e := range prerequisites {
		g[e[0]] = append(g[e[0]], e[1])
	}
	for i := range sp {
		sp[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		vis := make([]bool, n)
		dfs(i, i, vis, g, sp)
	}
	for _, e := range queries {
		res = append(res, sp[e[0]][e[1]])
	}
	return res
}
