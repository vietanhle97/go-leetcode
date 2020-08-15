package Find_K_th_Smallest_Pair_Distance_719

import "sort"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	m := len(nums)
	if m == 2 {
		return nums[1] - nums[0]
	}
	l, r := 0, nums[len(nums)-1]-nums[0]

	for l <= r {
		mid := l + (r-l)/2
		start, cnt := 0, 0
		for i, e := range nums {
			for start < len(nums) && nums[start]-e <= mid {
				start++
			}
			cnt += start - i - 1
		}
		if cnt >= k {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
