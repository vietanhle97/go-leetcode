package Find_K_Closest_Elements_658

import (
	"math"
	"sort"
)

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func findMaxSmaller(nums []int, l, r, x int) int {
	res := 0
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] <= x {
			res = max(res, mid)
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}

func findClosestElements(arr []int, k int, x int) []int {
	m := len(arr)
	if k == m {
		return arr
	}
	res := make([]int, 0)
	ls := findMaxSmaller(arr, 0, m-1, x)
	rs := ls + 1
	if ls == 0 {
		return arr[:k]
	}
	if ls == m-1 {
		return arr[m-k : m]
	}
	for len(res) < k {
		if ls < 0 {
			res = append(res, arr[rs])
			rs++
		} else if rs >= m {
			res = append(res, arr[ls])
			ls--
		} else {
			if x-arr[ls] <= arr[rs]-x {
				res = append(res, arr[ls])
				ls--
			} else {
				res = append(res, arr[rs])
				rs++
			}
		}
	}
	sort.Ints(res)
	return res
}
