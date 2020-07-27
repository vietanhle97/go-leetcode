package Minimum_Height_Trees_310

func findMinHeightTrees(n int, edges [][]int) []int {
	if len(edges) == 0 {
		return []int{0}
	}
	g := make([][]int, n)
	degree := make([]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
		degree[e[0]]++
		degree[e[1]]++
	}
	q := make([]int, 0)
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			q = append(q, i)
		}
	}
	for n > 2 {
		tmp := len(q)
		for i := 0; i < tmp; i++ {
			t := q[0]
			q = q[1:]
			n--
			for _, e := range g[t] {
				degree[e]--
				if degree[e] == 1 {
					q = append(q, e)
				}
			}
		}
	}
	return q
}
