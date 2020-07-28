package Redundant_Connection_II_685

func find(parent []int, x int) int {
	if parent[x] == 0 {
		parent[x] = x
	}
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

func findRedundantDirectedConnection(edges [][]int) []int {
	m := len(edges)
	g := make([]int, m+1)
	parent := make([]int, m+1)
	ans1 := make([]int, 2)
	ans2 := make([]int, 2)
	for _, ed := range edges {
		u, v := ed[0], ed[1]
		if g[v] > 0 {
			ans1 = []int{g[v], v}
			ans2 = ed
		}
		g[v] = u
	}

	for _, ed := range edges {
		if ans2[0] == ed[0] && ans2[1] == ed[1] {
			continue
		}
		u, v := ed[0], ed[1]
		pu, pv := find(parent, u), find(parent, v)
		if pu == pv {
			if ans1[0] != 0 && ans1[1] != 0 {
				return ans1
			}
			return ed
		}
		parent[pv] = pu
	}
	return ans2
}
