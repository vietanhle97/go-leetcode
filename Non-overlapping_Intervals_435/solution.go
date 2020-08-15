package Non_overlapping_Intervals_435

import "sort"

func eraseOverlapIntervals(A [][]int) int {
	if len(A) <= 1 {
		return 0
	}
	sort.Slice(A, func(i, j int) bool {
		if A[i][1] == A[j][1] {
			return A[i][0] < A[j][0]
		}
		return A[i][1] < A[j][1]
	})
	st := [][]int{A[0]}
	for i := 1; i < len(A); i++ {
		if A[i][0] < st[len(st)-1][1] {
			continue
		} else {
			st = append(st, A[i])
		}
	}
	return len(A) - len(st)
}
