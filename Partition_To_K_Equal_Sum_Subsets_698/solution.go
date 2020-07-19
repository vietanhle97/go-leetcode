package Partition_To_K_Equal_Sum_Subsets_698

import "sort"

func countNum(nums []int) bool {
	cnt := map[int]bool{}
	res := 0
	for _, e := range nums {
		if _, ok := cnt[e]; !ok {
			res++
			cnt[e] = true
			if len(cnt) > 1 {
				return false
			}

		}
	}
	return true
}

func sum(nums []int) int {
	res := 0
	for _, e := range nums {
		res += e
	}
	return res
}

func dfs(i, target int, MAX int, nums, sum []int, k int) bool {
	if i == MAX {
		return countNum(sum)
	}
	for j := 0; j < k; j++ {
		sum[j] += nums[i]
		if sum[j] <= target && dfs(i+1, target, MAX, nums, sum, k) {
			return true
		}
		sum[j] -= nums[i]
		if sum[j] == 0 {
			break
		}
	}
	return false
}

func canPartitionKSubsets(nums []int, k int) bool {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	m := len(nums)
	if m == 0 {
		return false
	}
	SUM := sum(nums)
	if SUM%k != 0 {
		return false
	}

	target := SUM / k
	sub := make([]int, k)
	return dfs(0, target, m, nums, sub, k)
}
