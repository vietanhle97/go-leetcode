package Combination_Sum_39

import (
	"sort"
	"strconv"
)

func backtrack(start int, size int, target int, nums []int, path []int, sum_ int, m *map[string]bool, res *[][]int) {
	if sum_ == target {
		tmp := ""
		for _, e := range path {
			tmp += strconv.Itoa(e)
		}
		if _, ok := (*m)[tmp]; !ok {
			*res = append(*res, append([]int{}, path...))
			(*m)[tmp] = true
		}

		return
	}
	for i := start; i < size; i++ {
		if sum_+nums[i] <= target {
			backtrack(i, size, target, nums, append(path, nums[i]), sum_+nums[i], m, res)
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	m := map[string]bool{}
	sort.Ints(candidates)
	backtrack(0, len(candidates), target, candidates, []int{}, 0, &m, &res)
	return res
}
