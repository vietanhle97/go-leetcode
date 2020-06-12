package Number_Of_Islands_200

func isValid(i, j, m, n int, grid [][]byte) bool {
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == byte(48) {
		return false
	}
	return true
}

func find(parent []int, i int) int {
	if parent[i] < 0 {
		return i
	}
	parent[i] = find(parent, parent[i])
	return parent[i]
}

func union(parent []int, i, j int, count *int) {
	Pi, Pj := find(parent, i), find(parent, j)
	if Pi == Pj {
		return
	}
	if Pi > Pj {
		Pi, Pj = Pj, Pi
	}
	parent[Pi] += parent[Pj]
	parent[Pj] = Pi
	*count -= 1
}

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	count := 0
	parent := make([]int, m*n)
	dir := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for i := range parent {
		if grid[i/n][i%n] == byte(49) {
			parent[i] = -1
			count += 1
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == byte(49) {
				for _, k := range dir {
					if isValid(i+k[0], j+k[1], m, n, grid) {
						union(parent, n*i+j, n*(i+k[0])+j+k[1], &count)
					}
				}
			}
		}
	}
	return count
}
