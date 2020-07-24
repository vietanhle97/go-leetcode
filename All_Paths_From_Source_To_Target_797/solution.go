package All_Paths_From_Source_To_Target_797

func dfs(v, target int, g map[int][]int, res *[][]int, vis []bool, path []int) {
	if v == target {
		path = append(path, v)
		*res = append(*res, append([]int{}, path...))
		return
	}
	vis[v] = true
	for _, e := range g[v] {
		if !vis[e] {
			dfs(e, target, g, res, vis, append(path, v))
		}
	}
	vis[v] = false
}

func allPathsSourceTarget(graph [][]int) [][]int {
	g := map[int][]int{}
	MAX := 0
	for i, e := range graph {
		g[i] = e
		for _, v := range e {
			if v > MAX {
				MAX = v
			}
		}
	}
	vis := make([]bool, MAX+1)
	res := make([][]int, 0)
	dfs(0, MAX, g, &res, vis, []int{})
	return res
}
