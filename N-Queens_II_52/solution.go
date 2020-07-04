package N_Queens_II_52

import "math"

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func valid(r, c int, path []int) bool {
	for i := range path {
		if abs(i-c) == abs(r-path[i]) {
			return false
		}
	}
	return true
}

func concat(b, c []int) []int {
	m, n := len(b), len(c)
	a := make([]int, m+n)
	for i := 0; i < m+n; i++ {
		if i < m {
			a[i] = b[i]
		} else {
			a[i] = c[i-m]
		}
	}
	return a
}

func backtrack(nums, path []int, max int, res *int) {
	if len(nums) <= 0 {
		if len(path) == max {
			*res++
		}
		return
	}
	for i, e := range nums {
		if len(path) == 0 {
			backtrack(concat(nums[:i], nums[i+1:]), append(path, e), max, res)
		} else {
			if valid(nums[i], len(path), path) {
				backtrack(concat(nums[:i], nums[i+1:]), append(path, e), max, res)
			}
		}
	}
}

func totalNQueens(n int) int {
	res := 0
	if n == 1 {
		return 1
	}
	if n <= 3 {
		return res
	}
	nums := make([]int, n)
	for i := 1; i < n+1; i++ {
		nums[i-1] = i
	}
	backtrack(nums, []int{}, n, &res)
	return res
}
