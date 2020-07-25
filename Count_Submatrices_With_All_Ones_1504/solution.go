package Count_Submatrices_With_All_Ones_1504

func numSubmat(A [][]int) int {
	R, C := len(A), len(A[0])
	res := 0
	for r, row := range A {
		for c, v := range row {
			if c > 0 && v > 0 {
				A[r][c] += A[r][c-1]
			}
		}
	}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			tmp := A[r][c]
			for r2 := r; r2 < R; r2++ {
				if A[r2][c] < tmp {
					tmp = A[r2][c]
				}
				res += tmp
			}
		}
	}
	return res
}
