package Maximum_Side_Length_Of_A_Square_With_Sum_Less_Than_Or_Equal_To_Threshold_1292

import (
	"math"
)

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
func calculateSum(mat [][]int, m, n int) [][]int {
	prefixSum := make([][]int, m+1)
	prefixSum[0] = make([]int, n+1)
	for i := 1; i < m+1; i++ {
		prefixSum[i] = make([]int, n+1)
		tmp := 0
		for j := 1; j < n+1; j++ {
			tmp += mat[i-1][j-1]
			prefixSum[i][j] = tmp
		}
	}
	return prefixSum
}

func isValid(prefixSum [][]int, side, m, n, threshold int) bool {
	for i := 1; i <= m-side+1; i++ {
		for j := side; j < n+1; j++ {
			sum := 0
			for k := i; k < i+side; k++ {
				sum += prefixSum[k][j] - prefixSum[k][j-side]
			}
			if sum <= threshold {
				return true
			}
		}
	}
	return false
}

func maxSideLength(mat [][]int, threshold int) int {
	m := len(mat)

	if m == 1 {
		return 1
	}
	n := len(mat[0])
	if n == 1 {
		return 1
	}
	prefixSum := calculateSum(mat, m, n)
	l, r := 1, min(m, n)
	res := 0
	for l <= r {
		mid := l + (r-l)/2
		if isValid(prefixSum, mid, m, n, threshold) {
			if res == 0 {
				res = mid
			}
			res = max(res, mid)
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}
