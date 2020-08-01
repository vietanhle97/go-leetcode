package H_Index_274

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)
	m := len(citations)
	lo, hi := 0, m
	for lo < hi {
		mid := lo + (hi-lo)/2
		if citations[mid] >= m-mid {
			if mid == 0 || citations[mid-1] < m-(mid-1) {
				return m - mid
			}
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return 0
}
