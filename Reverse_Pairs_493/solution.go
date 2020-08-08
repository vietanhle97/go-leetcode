package Reverse_Pairs_493

import "sort"

var bit []int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func count(i int) int {
	s := 0
	for i < len(bit) {
		s += bit[i]
		i += i & (-i)
	}
	return s
}

func update(i, val int) {
	for i > 0 {
		bit[i] += val
		i -= i & (-i)
	}
}

func lowBound(A []int, v int) int {
	l, r := 0, len(A)
	for l < r {
		mid := l + (r-l)/2
		if v <= A[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l == len(A) {
		if A[l-1] >= v {
			return l - 1
		}
		return l
	}
	return l
}

func reversePairs(A []int) int {
	m, res := len(A), 0
	tmp := make([]int, m)
	for i, e := range A {
		tmp[i] = e
	}
	sort.Ints(tmp)
	bit = make([]int, m+1)
	for i := 0; i < m; i++ {
		res += count(lowBound(tmp, 2*A[i]+1) + 1)
		update(lowBound(tmp, A[i])+1, 1)
	}
	return res
}
