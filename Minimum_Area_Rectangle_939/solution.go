package Minimum_Area_Rectangle_939

import "sort"

type node struct {
	y1 int
	y2 int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func minAreaRect(points [][]int) int {
	cols := make([]int, 0)
	m := map[int][]int{}
	y := map[node]int{}
	res := 1<<31 - 1
	for _, e := range points {
		if _, ok := m[e[0]]; !ok {
			cols = append(cols, e[0])
		}
		m[e[0]] = append(m[e[0]], e[1])
	}
	sort.Ints(cols)
	for _, x := range cols {
		for j, y1 := range m[x] {
			for i := 0; i < j; i++ {
				y2 := m[x][i]
				ny1, ny2 := max(y1, y2), min(y1, y2)
				if _, ok := y[node{ny1, ny2}]; ok {
					res = min(res, (ny1-ny2)*abs(x-y[node{ny1, ny2}]))
				}
				y[node{ny1, ny2}] = x
			}
		}
	}
	if res == 1<<31-1 {
		return 0
	}
	return res
}
