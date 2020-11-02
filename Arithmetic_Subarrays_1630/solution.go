package Arithmetic_Subarrays_1630

import "sort"

func isArith(a []int) bool {
	if len(a) < 2 {
		return false
	}
	if len(a) == 2 {
		return true
	}
	sort.Ints(a)
	prev := a[1] - a[0]
	for i := 1; i < len(a)-1; i++ {
		if a[i+1]-a[i] != prev {
			return false
		}
	}
	return true
}
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	res := make([]bool, 0)
	for i := range l {
		tmp := append([]int{}, nums[l[i]:r[i]+1]...)
		res = append(res, isArith(tmp))
	}
	return res
}
