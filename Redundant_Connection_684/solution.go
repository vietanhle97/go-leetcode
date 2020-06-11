package Redundant_Connection_684

func find(parent []int, i int) int {
	if parent[i] == -1 {
		return i
	}
	return find(parent, parent[i])
}

func union(parent []int, x, y int) (bool, []int) {
	xSet := find(parent, x)
	ySet := find(parent, y)
	if xSet == ySet {
		return true, []int{x, y}
	}
	parent[xSet] = ySet
	return false, nil
}

func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	for i, _ := range parent {
		parent[i] = -1
	}
	for _, e := range edges {
		b, res := union(parent, e[0], e[1])
		if b {
			return res
		}
	}
	return []int{}
}
