package Largest_Number_179

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	m := len(nums)
	N := make([]string, 0)
	for i := range nums {
		N = append(N, strconv.Itoa(nums[i]))
	}
	sort.Slice(N, func(i, j int) bool {
		return N[i]+N[j] < N[j]+N[i]
	})
	res := ""
	for i := m - 1; i > -1; i-- {
		if i == 0 {
			res += N[i]
			break
		}
		if len(res) == 0 && N[i] == "0" {
			continue
		}
		res += N[i]
	}
	return res
}
