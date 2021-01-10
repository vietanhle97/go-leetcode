package Construct_The_Lexicographically_Largest_Valid_Sequence_1718

func cmp(a, b []int) bool {
	for i := range a {
		if a[i] > b[i] {
			return true
		}
	}
	return false
}

func backtrack(cur, n, l int, A, path []int, res *[]int, exist *bool) {
	if *exist {
		return
	}
	if cur >= l {
		if cmp(path, *res) {
			for i := range path {
				(*res)[i] = path[i]
			}
			*exist = true
		}
		return
	}
	if path[cur] != 0 {
		backtrack(cur+1, n, l, A, path, res, exist)
		return
	}

	for i := n; i >= 1; i-- {
		b := false
		if A[i] > 1 && (cur+i >= l || path[cur+i] != 0) {
			continue
		}
		if A[i] > 0 {
			path[cur] = i
			A[i]--
			if i != 1 && A[i] > 0 && cur+i < l {
				b = true
				path[cur+i] = i
				A[i]--
			}
			backtrack(cur+1, n, l, A, path, res, exist)
			path[cur] = 0
			A[i]++
			if b {
				path[cur+i] = 0
				A[i]++
			}
		}
	}
}

func constructDistancedSequence(n int) []int {
	A := make([]int, n+1)
	A[1] = 1
	for i := 2; i <= n; i++ {
		A[i] = 2
	}
	res := make([]int, 2*n-1)
	exist := false
	backtrack(0, n, 2*n-1, A, make([]int, 2*n-1), &res, &exist)
	return res
}
