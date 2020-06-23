package Find_The_Smallest_Divisor_Given_A_Threshold_1283

import "math"

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func findMax(nums []int) int {
	max := nums[0]
	for _, e := range nums {
		if e > max {
			max = e
		}
	}
	return max
}

func isValid(nums []int, d, threshold int) bool {
	sum := 0
	for _, e := range nums {
		if e%d == 0 {
			sum += e / d
		} else {
			sum += e/d + 1
		}
	}
	return sum <= threshold
}

func smallestDivisor(nums []int, threshold int) int {
	l, r := 1, findMax(nums)
	res := r
	for l < r {
		mid := l + (r-l)/2
		if isValid(nums, mid, threshold) {
			res = min(res, mid)
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}
