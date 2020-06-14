package Contains_Duplicate_III_220

import "sort"

func findMinGreater(arr []int, n int) int {
	lo, hi := 0, len(arr)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if arr[mid] < n {
			lo = mid + 1
		} else if arr[mid] > n {
			hi = mid
		} else {
			return mid + 1
		}
	}
	return lo
}

func findMaxSmaller(arr []int, n int) int {
	lo, hi := 0, len(arr)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if arr[mid] < n {
			lo = mid + 1
		} else if arr[mid] > n {
			hi = mid
		} else {
			return mid - 1
		}
	}
	return hi
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	tmp := make([]int, 0)
	check := map[int]int{}
	for i, e := range nums {
		if len(tmp) == 0 {
			tmp = append(tmp, e)
			check[i] = e
		} else {
			sort.Ints(tmp)
			min := tmp[findMinGreater(tmp, e)]
			if e-min <= t {
				return true
			}
			max := tmp[findMaxSmaller(tmp, e)]
			if max-e <= t {
				return true
			}
			tmp = append(tmp, e)
			check[i] = e
			if len(tmp) > k {
				tmp = tmp[1:]
				delete(check, i-k)
			}
		}
	}
	return false
}
