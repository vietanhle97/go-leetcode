package Koko_Eating_Bananas_875

import "math"

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func findMax(piles []int) int {
	max := piles[0]
	for _, e := range piles {
		if e > max {
			max = e
		}
	}
	return max
}
func isValid(piles []int, H, K int) bool {
	cnt := 0
	for _, p := range piles {
		cnt += (p-1)/K + 1
	}
	return cnt <= H
}

func minEatingSpeed(piles []int, H int) int {
	r := findMax(piles)
	l := 1
	cnt := -1
	for l <= r {
		mid := l + (r-l)/2
		if isValid(piles, H, mid) {

			if cnt == -1 {
				cnt = mid
			}
			cnt = min(cnt, mid)
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return cnt
}
