package Range_Sum_Of_Sorted_Subarray_Sums_1508

import "sort"

func rangeSum(nums []int, n int, l int, r int) int {

	res := make([]int, 0)
	m := len(nums)
	MOD := int(1e9 + 7)
	for i, e1 := range nums {
		sum := e1
		res = append(res, e1)
		for j := i + 1; j < m; j++ {
			sum += nums[j]
			res = append(res, sum)
		}
	}
	sort.Ints(res)
	ans := 0
	for i := l - 1; i < r; i++ {
		ans += res[i] % MOD
	}
	return ans % MOD
}
