package Subsets_78

import "sort"

func backtrack(start, m int, nums, path []int, res *[][]int) {
	*res = append(*res, append([]int{}, path...))
	if start >= m {
		return
	}
	for i := start; i < m; i++ {
		backtrack(i+1, m, nums, append(path, nums[i]), res)
	}
}

func subsets(nums []int) [][]int {
	m := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)
	backtrack(0, m, nums, []int{}, &res)
	return res
}
