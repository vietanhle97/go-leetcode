package Heaters_475

import "sort"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func canCover(houses, heaters []int, v int) bool {
	m, n := len(heaters), len(houses)
	if heaters[0]-v > houses[0] {
		return false
	}
	j := 0
	for i := 0; i < m; i++ {
		for j < n {
			if heaters[i] > houses[j] {
				if heaters[i]-v <= houses[j] {
					j++
					if j == n {
						return true
					}
				} else {
					break
				}
			} else {
				if heaters[i]+v >= houses[j] {
					j++
					if j == n {
						return true
					}
				} else {
					break
				}
			}
		}
	}
	return false
}

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	l, r := 0, int(1e9)
	res := 1<<31 - 1
	for l <= r {
		mid := l + (r-l)/2
		if canCover(houses, heaters, mid) {
			res = min(mid, res)
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}
