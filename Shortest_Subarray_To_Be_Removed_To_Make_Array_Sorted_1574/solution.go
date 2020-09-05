package Shortest_Subarray_To_Be_Removed_To_Make_Array_Sorted_1574

func bs(l, r, v int, arr []int, first bool) int {
	if l > r {
		return -1
	}
	mid := l + (r-l)/2
	if !first {
		rs := bs(mid+1, r, v, arr, first)
		if rs != -1 {
			return rs
		}
		if arr[mid] <= v {
			return mid
		}
		return bs(l, mid-1, v, arr, first)
	} else {
		ls := bs(l, mid-1, v, arr, first)
		if ls != -1 {
			return ls
		}
		if arr[mid] >= v {
			return mid
		}
		return bs(mid+1, r, v, arr, first)
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	pref, suff := 0, n-1
	for i := 1; i < n; i++ {
		if arr[i] >= arr[pref] {
			pref = i
		} else {
			break
		}
	}
	if pref == n-1 {
		return 0
	}
	for i := n - 2; i >= 0; i-- {
		if arr[i] <= arr[suff] {
			suff = i
		} else {
			break
		}
	}
	l, r := bs(0, pref, arr[suff], arr, false), bs(suff, n-1, arr[pref], arr, true)
	if l == -1 && r == -1 {
		return min(n-1-pref, suff)
	} else if r == -1 {
		return min(n-1-pref, suff-l-1)
	} else if l == -1 {
		return min(suff, r-pref-1)
	}
	return min(suff-l-1, r-pref-1)
}
