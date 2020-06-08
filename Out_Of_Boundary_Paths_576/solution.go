package Out_Of_Boundary_Paths_576

import "math"

func findPaths(m int, n int, N int, x int, y int) int {
	res := make([][]int, m)
	MOD := int(math.Pow10(9)) + 7
	for i := range res {
		res[i] = make([]int, n)
	}
	res[x][y] = 1
	cnt := 0
	for move := 1; move <= N; move++ {
		tmp := make([][]int, m)
		for i := range tmp {
			tmp[i] = make([]int, n)
		}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {

				if i == m-1 {
					// fmt.Println(i,j,res[i][j])
					cnt = (cnt + res[i][j]) % MOD
				}
				if j == n-1 {
					// fmt.Println(i,j,res[i][j])
					cnt = (cnt + res[i][j]) % MOD
				}
				if i == 0 {
					cnt = (cnt + res[i][j]) % MOD
				}
				if j == 0 {
					cnt = (cnt + res[i][j]) % MOD
				}
				a, b, c, d := 0, 0, 0, 0
				if i > 0 {
					a = res[i-1][j]
				}
				if i < m-1 {
					b = res[i+1][j]
				}
				if j > 0 {
					c = res[i][j-1]
				}
				if j < n-1 {
					d = res[i][j+1]
				}
				tmp[i][j] = (a+b)%MOD + (c+d)%MOD
			}
		}
		res = tmp
	}
	return cnt
}
