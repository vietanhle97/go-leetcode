package Permutations_46

func dfs(nums []int, path []int, check map[int]bool, res *[][]int) {
	if len(path) == len(nums) {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for _, e := range nums {
		if val, _ := check[e]; !val {
			check[e] = true
			dfs(nums, append(path, e), check, res)
			check[e] = false
		}
	}
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	check := map[int]bool{}
	for _, e := range nums {
		check[e] = false
	}
	dfs(nums, []int{}, check, &res)
	return res
}
