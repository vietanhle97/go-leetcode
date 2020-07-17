package Smallest_Range_II_910

import "sort"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func smallestRangeII(A []int, K int) int {
	sort.Ints(A)

	res := A[len(A)-1] - A[0]
	for i := 0; i < len(A)-1; i++ {
		MIN := min(A[0]+K, A[i+1]-K)
		MAX := max(A[len(A)-1]-K, A[i]+K)
		res = min(res, MAX-MIN)
	}
	return res
}
