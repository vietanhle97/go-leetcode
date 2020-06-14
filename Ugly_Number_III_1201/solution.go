package Ugly_Number_III_1201

import (
	"math"
	"sort"
)

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func sum(mid int, arr []int) int {
	res := 0
	for _, e := range arr {
		res += mid / e
	}
	return res
}

func nthUglyNumber(n int, a int, b int, c int) int {

	nums := []int{a, b, c}
	sort.Ints(nums)
	nums2 := []int{lcm(a, b), lcm(b, c), lcm(c, a)}
	LCM := lcm(a, lcm(b, c))

	low, high := n, int(2*math.Pow10(9))
	for low < high {
		mid := low + (high-low)/2
		rank := sum(mid, nums) - sum(mid, nums2) + mid/LCM
		if rank < n {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
