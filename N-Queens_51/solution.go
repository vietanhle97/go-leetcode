package N_Queens_51

import (
	"math"
)

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

func ans(i, n int) string {
	res := ""
	for k := 0; k < i; k++ {
		res += "."
	}
	res += "Q"
	for k := i + 1; k < n; k++ {
		res += "."
	}
	return res
}
func backtrack(nums, path []int, max int, res *[][]string, m map[int]string) {
	if len(nums) == 0 {
		if len(path) == max {
			tmp := make([]string, 0)
			for _, e := range path {
				tmp = append(tmp, m[e])
			}
			*res = append(*res, tmp)
		}
		return
	}
	for i, e := range nums {
		if len(path) == 0 {
			backtrack(concat(nums[:i], nums[i+1:]), append(path, e), max, res, m)
		} else {
			if valid(nums[i], len(path), path) {
				backtrack(concat(nums[:i], nums[i+1:]), append(path, e), max, res, m)
			}
		}
	}
}

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	m := map[int]string{}
	for i := 1; i < n+1; i++ {
		m[i] = ans(i-1, n)
	}
	if n == 1 {
		return [][]string{{"Q"}}
	}
	if n <= 3 {
		return res
	}
	nums := make([]int, n)
	for i := 1; i < n+1; i++ {
		nums[i-1] = i
	}
	backtrack(nums, []int{}, n, &res, m)
	return res
}
