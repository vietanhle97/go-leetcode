package Number_Of_Nodes_In_The_Sub_Tree_With_The_Same_Label_1519

func dfs(v int, g map[int][]int, labels string, m *map[int]int, vis *map[int]bool) []int {
	res := make([]int, 26)
	res[labels[v]-'a']++
	(*m)[v] = res[labels[v]-'a']
	(*vis)[v] = true
	for _, k := range g[v] {
		if (*vis)[k] == false {
			tmp := dfs(k, g, labels, m, vis)
			for i := range res {
				res[i] += tmp[i]
			}
			(*m)[v] = res[labels[v]-'a']
		}
	}
	return res
}

func countSubTrees(n int, edges [][]int, labels string) []int {
	g := map[int][]int{}
	res := map[int]int{}
	vis := map[int]bool{}
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		if e[0] != 0 {
			g[e[1]] = append(g[e[1]], e[0])
		}
	}
	for i := 1; i < n; i++ {
		res[i] = 0
		vis[i] = false
	}
	cnt := make([]int, n)
	dfs(0, g, labels, &res, &vis)
	for k, v := range res {
		cnt[k] = v
	}
	return cnt
}
