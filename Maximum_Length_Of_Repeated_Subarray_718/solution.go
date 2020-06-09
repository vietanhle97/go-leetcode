package Maximum_Length_Of_Repeated_Subarray_718

import "math"

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func findLength(A []int, B []int) int {
	m := len(A)
	n := len(B)
	res := 0
	if m == 0 || n == 0 {
		return 0
	}
	table := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		table[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if A[i-1] == B[j-1] {
				table[i][j] = table[i-1][j-1] + 1
				res = max(res, table[i][j])
			} else {
				continue
			}
		}
	}
	return res
}
