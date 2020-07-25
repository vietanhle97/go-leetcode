package Maximal_Square_221

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maximalSquare(A [][]byte) int {
	R := len(A)
	if R == 0 {
		return 0
	}
	C := len(A[0])
	if C == 0 {
		return 0
	}
	B := make([][]int, R+1)
	for r := 0; r < R+1; r++ {
		B[r] = make([]int, C+1)
	}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			tmp := 0
			if A[r][c] == '1' {
				tmp = 1
			}
			B[r+1][c+1] = tmp + B[r+1][c] + B[r][c+1] - B[r][c]
		}
	}

	res := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			for r1 := r; r1 < R; r1++ {
				if r1-r+c+1 <= C {
					interior := B[r1+1][r1-r+c+1] - B[r1+1][c] - B[r][r1-r+c+1] + B[r][c]
					if interior == (r1-r+1)*(r1-r+1) {
						res = max(res, (r1-r+1)*(r1-r+1))
					}
				}

			}
		}
	}
	return res
}
