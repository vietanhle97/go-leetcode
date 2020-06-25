package Valid_Triangle_Number_611

import (
	"math"
	"sort"
)

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func findMaxSmaller(nums []int, l, r, target int) int {
	res := -1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			if res == -1 {
				res = mid
			}
			if res != -1 {
				res = max(res, mid)
			}
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}

func findMinLarger(nums []int, l, r, target int) int {
	res := -1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
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

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			rs := findMaxSmaller(nums, j+1, len(nums)-1, nums[i]+nums[j])
			if rs == -1 {
				continue
			}
			ls := findMinLarger(nums, j+1, len(nums)-1, nums[j]-nums[i])
			if ls == -1 {
				continue
			}
			tmp := rs - ls + 1
			if tmp > 0 {
				res += tmp
			}
		}
	}
	return res
}
