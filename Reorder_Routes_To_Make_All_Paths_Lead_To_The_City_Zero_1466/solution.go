package Reorder_Routes_To_Make_All_Paths_Lead_To_The_City_Zero_1466

func check(arr []int, v int) bool {
	for _, e := range arr {
		if v == e {
			return true
		}
	}
	return false
}
func dfs(v int, g map[int][]int, out map[int][]int, vis []bool, res *int) {
	vis[v] = true
	for _, e := range g[v] {
		if !vis[e] {
			if v != 0 && check(out[e], v) {
				*res++
			}
			dfs(e, g, out, vis, res)
		}
	}
}

func minReorder(n int, connections [][]int) int {
	g := map[int][]int{}
	out := map[int][]int{}
	res := 0
	for _, e := range connections {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
		out[e[1]] = append(out[e[1]], e[0])
		if e[0] == 0 {
			res++
		}
	}
	vis := make([]bool, n)
	dfs(0, g, out, vis, &res)
	return res
}
