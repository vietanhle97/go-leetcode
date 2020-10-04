package Evaluate_Division_399

func dfs(s, d int, g [][]float64, vis []bool, cur float64, res *float64) {
	if s == d {
		*res = cur
		return
	}
	vis[s] = true
	for i, e := range g[s] {
		if !vis[i] && e != -1.0 {
			dfs(i, d, g, vis, cur*e, res)
		}
	}
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	cnt := 0
	m := map[string]int{}
	check := map[string]bool{}
	for _, e := range equations {
		check[e[0]] = true
		check[e[1]] = true
		if _, ok := m[e[0]]; !ok {
			m[e[0]] = cnt
			cnt++
		}
		if _, ok := m[e[1]]; !ok {
			m[e[1]] = cnt
			cnt++
		}
	}
	g, ans := make([][]float64, len(check)), make([]float64, 0)
	for i := range g {
		g[i] = make([]float64, 26)
		for j := range g[i] {
			g[i][j] = -1.0
		}
	}
	for i, e := range equations {
		g[m[e[0]]][m[e[1]]] = values[i]
		g[m[e[1]]][m[e[0]]] = 1 / values[i]
	}
	for _, e := range queries {
		if _, ok := check[e[0]]; !ok {
			ans = append(ans, -1.0)
			continue
		}
		if _, ok := check[e[1]]; !ok {
			ans = append(ans, -1.0)
			continue
		}
		vis := make([]bool, 26)
		res := -1.0
		dfs(m[e[0]], m[e[1]], g, vis, 1.0, &res)
		ans = append(ans, res)
	}
	return ans
}
