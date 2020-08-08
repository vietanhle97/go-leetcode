package Count_Of_Range_Sum_327

import "sort"

var bit []int

func count(i int) int {
	s := 0
	for i > 0 {
		s += bit[i]
		i -= i & (-i)
	}
	return s
}

func update(i, v int) {
	for i < len(bit) {
		bit[i] += v
		i += i & (-i)
	}
}

func lowBound(A []int, v int) int {
	l, r := 0, len(A)-1
	for l < r {
		mid := (r + l) / 2
		if A[mid] < v {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if len(A) > 0 && A[l] >= v {
		return l
	}
	return -1
}

func upBound(A []int, v int) int {
	l, r := 0, len(A)-1
	for l < r {
		mid := l + 1 + (r-l)/2
		if A[mid] <= v {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if len(A) > 0 && A[l] <= v {
		return l
	}
	return -1
}

func countRangeSum(nums []int, l int, r int) int {
	res := 0
	sum, pref := []int{0}, []int{0}
	for _, e := range nums {
		sum = append(sum, e+sum[len(sum)-1])
		pref = append(pref, e+pref[len(pref)-1])
	}
	bit = make([]int, len(sum)+1)
	sort.Ints(pref)
	for _, e := range sum {
		lo, hi := lowBound(pref, e-r), upBound(pref, e-l)
		if lo != -1 && hi != -1 {
			res += count(hi+1) - count(lo)
		}
		update(1+lowBound(pref, e), 1)
	}
	return res
}
