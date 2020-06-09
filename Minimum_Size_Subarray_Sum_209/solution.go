package Minimum_Size_Subarray_Sum_209

import "math"

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func minSubArrayLen(s int, nums []int) int {
	res := len(nums) + 1
	sum, l := 0, 0

	for r, v := range nums {
		sum += v
		for sum >= s {
			res = min(res, r-l+1)
			sum -= nums[l]
			l += 1
		}
	}
	if res <= len(nums) {
		return res
	}
	return 0
}
