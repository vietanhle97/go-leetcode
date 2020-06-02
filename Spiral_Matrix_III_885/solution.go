package Spiral_Matrix_III_885

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	res := [][]int{{r0, c0}}
	if R*C == 1 {
		return res
	}
	for i := 1; i < 2*(R+C); i += 2 {
		dir := [][]int{{0, 1, i}, {1, 0, i}, {0, -1, i + 1}, {-1, 0, i + 1}}
		for _, j := range dir {
			for k := 0; k < j[2]; k++ {
				r0 += j[0]
				c0 += j[1]
				if 0 <= r0 && r0 < R && 0 <= c0 && c0 < C {
					res = append(res, []int{r0, c0})
					if len(res) == R*C {
						return res
					}
				}
			}
		}
	}
	return res
}
