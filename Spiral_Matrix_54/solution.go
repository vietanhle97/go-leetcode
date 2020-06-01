package Spiral_Matrix_54

func isValid(x, y, m, n int, check map[int]bool, matrix [][]int) bool {
	if x < 0 || x >= m || y < 0 || y >= n {
		return false
	}
	if ok, _ := check[matrix[x][y]]; ok {
		return false
	}
	return true
}

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	if n == 0 {
		return []int{}
	}
	res := make([]int, 0)
	check := map[int]bool{}
	cnt := 0
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	cntDir := 0
	x, y := 0, 0
	for cnt < m*n {
		res = append(res, matrix[x][y])
		cnt += 1
		check[matrix[x][y]] = true
		if !isValid(x+dir[cntDir][0], y+dir[cntDir][1], m, n, check, matrix) {
			cntDir = (cntDir + 1) % 4
		}
		x += dir[cntDir][0]
		y += dir[cntDir][1]
	}
	return res
}
