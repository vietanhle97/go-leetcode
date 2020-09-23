package Min_Cost_To_Connect_All_Points_1584

import "sort"

var parent, rank []int

type Edge struct {
	u int
	v int
	w int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func makeEdge(P [][]int, i, j int) Edge {
	w := abs(P[i][0]-P[j][0]) + abs(P[i][1]-P[j][1])
	return Edge{i, j, w}
}

func find(x int) int {
	if parent[x] == x {
		return x
	}
	return find(parent[x])
}

func union(x, y int) {
	px, py := find(x), find(y)
	if rank[px] < rank[py] {
		parent[px] = py
	} else if rank[py] < rank[px] {
		parent[py] = px
	} else {
		parent[py] = px
		rank[px]++
	}
}

func minCostConnectPoints(P [][]int) int {
	m := len(P)
	parent, rank = make([]int, m), make([]int, m)
	for i := range P {
		parent[i] = i
	}
	E := make([]Edge, 0)
	for i := 0; i < m-1; i++ {
		for j := i + 1; j < m; j++ {
			E = append(E, makeEdge(P, i, j))
		}
	}
	sort.Slice(E, func(i, j int) bool {
		return E[i].w < E[j].w
	})

	cnt, res, i := 0, 0, 0
	for cnt < m-1 {
		e := E[i]
		i++
		x, y := find(e.u), find(e.v)
		if x != y {
			cnt++
			union(x, y)
			res += e.w
		}
	}
	return res
}
