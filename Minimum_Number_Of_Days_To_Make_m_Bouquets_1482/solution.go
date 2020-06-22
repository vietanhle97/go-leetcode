package Minimum_Number_Of_Days_To_Make_m_Bouquets_1482

import "math"

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func isValid(bloomDay []int, m, k, mid int) bool {
	cnt, s := 0, 0
	for i := range bloomDay {
		if bloomDay[i] <= mid {
			s++
		} else {
			s = 0
		}
		if s == k {
			s = 0
			cnt++
		}
		if cnt == m {
			return true
		}
	}
	return false
}

func minDays(bloomDay []int, m int, k int) int {
	if len(bloomDay) == 0 || m == 0 || k == 0 {
		return 0
	}
	if m*k > len(bloomDay) {
		return -1
	}

	l, r := 1000000001, 0
	for _, e := range bloomDay {
		l = min(l, e)
		r = max(r, e)
	}

	for l <= r {
		mid := l + (r-l)/2
		if isValid(bloomDay, m, k, mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
