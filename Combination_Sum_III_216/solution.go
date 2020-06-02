package Combination_Sum_III_216

func backtrack(start int, target int, size int, path []int, res *[][]int) {
	if target == 0 && size == 0 {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := start; i < 10; i++ {
		if i <= target {
			backtrack(i+1, target-i, size-1, append(path, i), res)
		}
	}
}

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	backtrack(1, n, k, []int{}, &res)
	return res
}
