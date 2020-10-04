package K_diff_Pairs_In_An_Array_532

import "sort"

func bs(l, r, v, k int, arr []int) bool {
	if l > r {
		return false
	}
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid]-v == k {
			return true
		}
		if arr[mid]-v < k {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

func findPairs(nums []int, k int) int {
	sort.Ints(nums)
	m, l := map[int]bool{}, len(nums)-1
	res := 0
	for i, e := range nums {
		if _, ok := m[e]; !ok && bs(i+1, l, e, k, nums) {
			res++
			m[e] = true
		}
	}
	return res
}
