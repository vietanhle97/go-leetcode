package Minimum_Time_To_Collect_All_Apples_In_A_Tree_1443

type edge struct {
	u int
	v int
}

func dfs(v int, g map[int][]int, hasApple []bool, path []edge, res *map[edge]bool, vis *map[int]bool) {
	(*vis)[v] = true
	if hasApple[v] {
		for _, e := range path {
			if _, ok := (*res)[e]; !ok {
				(*res)[e] = true
			}
		}
	}
	for _, e := range g[v] {
		if (*vis)[e] == false {
			dfs(e, g, hasApple, append(path, edge{v, e}), res, vis)
		}
	}
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	g := map[int][]int{}
	res := map[edge]bool{}
	vis := map[int]bool{}
	for i := 0; i < n; i++ {
		g[i] = make([]int, 0)
		vis[i] = false
	}
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	dfs(0, g, hasApple, []edge{}, &res, &vis)
	return 2 * len(res)
}
