package Permutation_Sequence_60

import (
	"strconv"
)

func factorial(n int) int {
	res := 1
	for n > 0 {
		res *= n
		n--
	}
	return res
}

func getPermutation(n int, k int) string {
	nums := make([]int, n)
	for i, _ := range nums {
		nums[i] = i + 1
	}
	res := ""
	for len(res) < n {
		fac := factorial(len(nums) - 1)
		ind := k / fac
		r := k % fac
		if ind > 0 && r == 0 {
			ind -= 1
		}
		res += strconv.Itoa(nums[ind])
		k -= fac * ind

		nums = append(nums[:ind], nums[ind+1:]...)
	}
	return res
}
