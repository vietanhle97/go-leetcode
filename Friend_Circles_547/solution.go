package Friend_Circles_547

func find(parent []int, i int) int {
	for parent[i] != i {
		parent[i], i = parent[parent[i]], parent[i]
	}
	return i
}

func union(parent []int, x, y int) {
	xSet := find(parent, x)
	ySet := find(parent, y)
	if xSet < ySet {
		parent[ySet] = xSet
	} else if xSet > ySet {
		parent[xSet] = ySet
	}

}

func findCircleNum(M [][]int) int {
	m := len(M)
	parent := make([]int, len(M))
	check := map[int]interface{}{}
	for i := range parent {
		parent[i] = i
	}

	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			if M[i][j] == 1 {
				union(parent, i, j)
			}
		}
	}

	for _, e := range parent {
		check[find(parent, e)] = nil
	}
	return len(check)
}
