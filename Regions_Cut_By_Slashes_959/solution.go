package Regions_Cut_By_Slashes_959

func find(parent *[]int, x int) int {
	if (*parent)[x] != x {
		(*parent)[x] = find(parent, (*parent)[x])
	}
	return (*parent)[x]
}

func union(parent *[]int, x, y int) {
	xSet := find(parent, x)
	ySet := find(parent, y)
	(*parent)[xSet] = ySet
}

func regionsBySlashes(grid []string) int {
	N := len(grid)
	parent := make([]int, 4*N*N)
	res := 0
	for i := 0; i < 4*N*N; i++ {
		parent[i] = i
	}
	// fmt.Println(parent)
	for r, row := range grid {
		for c, col := range row {
			root := 4 * (r*N + c)
			if byte(col) == byte(32) || byte(col) == byte(47) {
				union(&parent, root+0, root+1)
				union(&parent, root+2, root+3)
			}
			if byte(col) == byte(32) || byte(col) == byte(92) {
				union(&parent, root+0, root+2)
				union(&parent, root+1, root+3)
			}

			if r+1 < N {
				union(&parent, root+3, root+4*N)
			}
			if r-1 >= 0 {
				union(&parent, root, root-4*N+3)
			}

			if c+1 < N {
				union(&parent, root+2, root+4+1)
			}
			if c-1 >= 0 {
				union(&parent, root+1, root-4+2)
			}
		}
	}
	for i := range parent {
		if find(&parent, i) == i {
			res++
		}
	}
	return res
}
