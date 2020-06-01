package Spiral_Matrix_II_59

func isValid(x, y, n int, matrix [][]int) bool {
	return x >= 0 && x < n && y >= 0 && y < n && matrix[x][y] == 0
}

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	cnt := 1
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	cntDir := 0
	x, y := 0, 0
	for cnt < n*n+1 {
		res[x][y] = cnt
		cnt += 1
		if !isValid(x+dir[cntDir][0], y+dir[cntDir][1], n, res) {
			cntDir = (cntDir + 1) % 4
		}
		x += dir[cntDir][0]
		y += dir[cntDir][1]
	}
	return res
}
