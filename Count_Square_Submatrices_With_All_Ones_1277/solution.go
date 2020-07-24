package Count_Square_Submatrices_With_All_Ones_1277

func countSquares(A [][]int) int {
	R, C := len(A), len(A[0])

	B := make([][]int, R+1)
	for r := 0; r < R+1; r++ {
		B[r] = make([]int, C+1)
	}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			B[r+1][c+1] = A[r][c] + B[r+1][c] + B[r][c+1] - B[r][c]
		}
	}

	res := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			for r1 := r; r1 < R; r1++ {
				if r1-r+c+1 <= C {
					interior := B[r1+1][r1-r+c+1] - B[r1+1][c] - B[r][r1-r+c+1] + B[r][c]
					if interior == (r1-r+1)*(r1-r+1) {
						res++
					}
				}

			}
		}
	}
	return res
}
