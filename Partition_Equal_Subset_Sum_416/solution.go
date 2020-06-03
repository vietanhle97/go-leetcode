package Partition_Equal_Subset_Sum_416

import "sort"

func sum(nums []int) int {
	res := 0
	for _, e := range nums {
		res += e
	}
	return res
}

func dfs(start, sum, max int, nums []int) bool {
	if sum == max/2 {
		return true
	} else if start >= len(nums) || sum > max/2 {
		return false
	}
	if nums[start] > max/2 {
		return false
	}
	res := false
	for i := start; i < len(nums); i++ {
		res = dfs(i+1, sum+nums[i], max, nums)
		if res {
			return res
		}
	}
	return false
}

func canPartition(nums []int) bool {
	s := sum(nums)
	if s%2 != 0 {
		return false
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return dfs(0, 0, s, nums)
}
