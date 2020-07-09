package Minimum_Cost_Tree_From_Leaf_Values_1130

import "math"

const MaxInt = 1<<31 - 1

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func findMinInd(nums []int) int {
	ind := 0
	for i := range nums {
		if nums[i] < nums[ind] {
			ind = i
		}
	}
	return ind
}

func mctFromLeafValues(nums []int) int {
	res := 0
	for len(nums) > 1 {
		minInd := findMinInd(nums)

		ls, rs := MaxInt, MaxInt

		if minInd > 0 {
			ls = nums[minInd-1]
		}

		if minInd < len(nums)-1 {
			rs = nums[minInd+1]
		}

		res += nums[minInd] * min(ls, rs)
		nums = append(nums[:minInd], nums[minInd+1:]...)
	}
	return res
}
