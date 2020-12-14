package Stone_Game_VI_1686

import "sort"

func stoneGameVI(A []int, B []int) int {
	m := len(A)
	sums := make([][]int, m)
	for i := 0; i < m; i++ {
		sums[i] = []int{A[i] + B[i], A[i], B[i]}
	}
	sort.Slice(sums, func(i, j int) bool {
		return sums[i][0] > sums[j][0]
	})
	a, b := 0, 0
	for i, e := range sums {
		if i%2 == 0 {
			a += e[1]
		} else {
			b += e[2]
		}
	}
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}
