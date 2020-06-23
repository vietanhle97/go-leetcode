package Capacity_To_Ship_Packages_Within_D_Days_1011

import "math"

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func isValid(weights []int, D, W int) bool {
	cnt := 0
	cur := 0
	for i, e := range weights {
		if W < e {
			return false
		}
		if cur+e <= W {
			cur += e
		} else {
			cnt++
			cur = e
		}
		if i == len(weights)-1 {
			cnt++
		}
	}
	return cnt <= D
}

func findSum(weights []int) int {
	sum := 0
	for _, e := range weights {
		sum += e
	}
	return sum
}

func shipWithinDays(weights []int, D int) int {
	r := findSum(weights)
	l := 1
	res := -1
	for l <= r {
		mid := l + (r-l)/2
		if isValid(weights, D, mid) {
			if res == -1 {
				res = mid
			}
			res = min(res, mid)
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}
