package Magnetic_Force_Between_Two_Balls_1552

import "sort"

func maxDistance(P []int, k int) int {
	sort.Ints(P)
	l, r := 1, int(1e9)
	for l <= r {
		mid, last, cnt := l+(r-l)/2, -int(1e9), 0
		for _, e := range P {
			if e-last >= mid {
				last = e
				cnt++
			}
		}
		if cnt >= k {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l - 1
}
