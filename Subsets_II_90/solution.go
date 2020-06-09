package Subsets_II_90

import (
	"sort"
	"strconv"
)

func convert(path []int) string {
	res := ""
	for i := range path {
		res += strconv.Itoa(path[i])
	}
	return res
}

func backtrack(start, m int, nums, path []int, res *[][]int, check *map[string]bool) {
	tmp := convert(path)
	if _, ok := (*check)[tmp]; !ok {
		*res = append(*res, append([]int{}, path...))
		(*check)[tmp] = true
	}
	if start >= m {
		return
	}
	for i := start; i < m; i++ {
		backtrack(i+1, m, nums, append(path, nums[i]), res, check)
	}

}

func subsetsWithDup(nums []int) [][]int {
	m := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)
	check := map[string]bool{}
	backtrack(0, m, nums, []int{}, &res, &check)
	return res
}
