package Minimum_Falling_Path_Sum_II_1289

import "math"

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func minFallingPathSum(A [][]int) int {
	m := len(A)
	n := len(A[0])
	ans := 20001
	for i := 1; i < m; i++ {
		ls := make([]int, n)
		rs := make([]int, n)
		ls[0] = A[i-1][0]
		for l := 1; l < n; l++ {
			ls[l] = min(ls[l-1], A[i-1][l])
		}
		rs[n-1] = A[i-1][n-1]
		for r := n - 2; r > -1; r-- {
			rs[r] = min(rs[r+1], A[i-1][r])
		}

		for j := 0; j < n; j++ {
			left, right := 20000, 20000
			if j-1 >= 0 {
				left = ls[j-1]
			}
			if j+1 < m {
				right = rs[j+1]
			}
			A[i][j] += min(left, right)
		}
	}
	for i := 0; i < n; i++ {
		ans = min(ans, A[m-1][i])
	}
	return ans
}
