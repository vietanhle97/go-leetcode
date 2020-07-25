package Critical_Connections_In_A_Network_1192

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func dfs(g [][]int, t, u, p int, low []int, res *[][]int) {
	low[u] = t
	// fmt.Println(low)
	for _, v := range g[u] {
		if p != v {
			if low[v] == -1 {
				dfs(g, t+1, v, u, low, res)
			}
			if low[v] > t {
				*res = append(*res, []int{u, v})
			} else {
				low[u] = min(low[u], low[v])
			}
		}
	}
}
func criticalConnections(n int, connections [][]int) [][]int {
	g := make([][]int, n)
	for _, e := range connections {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	low := make([]int, n)
	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		low[i] = -1
	}
	dfs(g, 1, 0, -1, low, &res)
	return res
}
