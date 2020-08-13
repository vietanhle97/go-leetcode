package As_Far_From_Land_As_Possible_1162

func valid(x, y, N int, A [][]int) bool {
	return x >= 0 && y >= 0 && x < N && y < N && A[x][y] != -1
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

func addVertex(x, y, N int, q *[][]int, dist [][]int) {
	if x+1 < N && dist[x+1][y] == -1 {
		*q = append(*q, []int{x + 1, y})
	}
	if x-1 >= 0 && dist[x-1][y] == -1 {
		*q = append(*q, []int{x - 1, y})
	}
	if y+1 < N && dist[x][y+1] == -1 {
		*q = append(*q, []int{x, y + 1})
	}
	if y-1 >= 0 && dist[x][y-1] == -1 {
		*q = append(*q, []int{x, y - 1})
	}
}

func maxDistance(A [][]int) int {
	x, y := []int{0, 0, 1, -1}, []int{1, -1, 0, 0}
	N, res := len(A), -1
	q, dist := make([][]int, 0), make([][]int, N)
	for i := range dist {
		dist[i] = make([]int, N)
		for j := range dist[i] {
			dist[i][j] = -1
			if A[i][j] == 1 {
				q = append(q, []int{i, j})
			}
		}
	}

	for len(q) > 0 {
		X, Y := q[0][0], q[0][1]
		q = q[1:]
		if A[X][Y] == 1 {
			dist[X][Y] = 0
			addVertex(X, Y, N, &q, dist)
		} else {
			tmp := N + 1
			for i := 0; i < 4; i++ {
				nX, nY := X+x[i], Y+y[i]
				if valid(nX, nY, N, dist) {
					tmp = min(tmp, dist[nX][nY])
				}
			}
			if dist[X][Y] == -1 {
				dist[X][Y] = tmp + 1
			}
			dist[X][Y] = min(dist[X][Y], tmp+1)
			res = max(dist[X][Y], res)
			addVertex(X, Y, N, &q, dist)
		}
	}
	return res
}
