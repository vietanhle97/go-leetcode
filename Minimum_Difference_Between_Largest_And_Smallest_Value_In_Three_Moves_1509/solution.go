package Minimum_Difference_Between_Largest_And_Smallest_Value_In_Three_Moves_1509

import "sort"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minDifference(nums []int) int {
	if len(nums) <= 3 {
		return 0
	}
	sort.Ints(nums)
	res := 1<<31 - 1
	MIN := nums[0:3]
	MAX := nums[len(nums)-3 : len(nums)]
	res = min(res, abs(MAX[2]-nums[3]))
	res = min(res, abs(MAX[1]-MIN[2]))
	res = min(res, abs(MAX[0]-MIN[1]))
	res = min(res, abs(nums[len(nums)-4]-MIN[0]))

	return res
}
