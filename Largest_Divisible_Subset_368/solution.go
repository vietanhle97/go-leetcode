package Largest_Divisible_Subset_368

import "sort"

type pair struct {
	prev int
	val  int
	cnt  int
}

func max(a, b pair) pair {
	if a.cnt > b.cnt {
		return a
	}
	return b
}

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	m := len(nums)
	table := make([]pair, 0)
	res := make([]int, 0)
	MAX := pair{}
	if m <= 1 {
		return nums
	}
	for i := 0; i < m; i++ {
		table = append(table, pair{-1, nums[i], 1})
		MAX = max(MAX, table[i])
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				tmp := pair{j, nums[i], table[j].cnt + 1}
				table[i] = max(table[i], tmp)
				MAX = max(MAX, table[i])
			}
		}
	}
	for MAX.prev != -1 {
		res = append(res, MAX.val)
		MAX = table[MAX.prev]
	}
	res = append(res, MAX.val)
	return res
}
