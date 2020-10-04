package Remove_Covered_Intervals_1288

import "sort"

func removeCoveredIntervals(A [][]int) int {
	sort.Slice(A, func(i, j int) bool {
		if A[i][0] == A[j][0] {
			return A[i][1] > A[j][1]
		}
		return A[i][0] < A[j][0]
	})
	st := make([][]int, 0)
	for _, e := range A {
		if len(st) == 0 {
			st = append(st, e)
		} else {
			if e[0] <= st[len(st)-1][1] && e[1] <= st[len(st)-1][1] {
				continue
			} else {
				st = append(st, e)
			}
		}
	}
	return len(st)
}
