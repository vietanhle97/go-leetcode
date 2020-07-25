package Maximal_Rectangle_85

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maximalRectangle(A [][]byte) int {
	R := len(A)
	if R == 0 {
		return 0
	}
	C := len(A[0])
	if C == 0 {
		return 0
	}
	res := 0

	for r, row := range A {
		for c, v := range row {
			if c > 0 && int(v) > 48 {
				A[r][c] += byte(int(A[r][c-1]) - 48)
			}
		}
	}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			tmp := int(A[r][c]) - 48
			for r2 := r; r2 < R; r2++ {
				if int(A[r2][c])-48 < tmp {
					tmp = int(A[r2][c]) - 48
				}
				res = max(res, tmp*(r2-r+1))
			}
		}
	}
	// fmt.Println(A)
	return res
}
