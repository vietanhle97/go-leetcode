package Permutations_II_47

func equal(arr1 []int, arr2 []int) bool {
	for i, _ := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func isExist(res *[][]int, path []int) bool {
	if len(*res) == 0 {
		return false
	}
	for _, e := range *res {
		if equal(path, e) {
			return true
		}
	}
	return false
}

func dfs(nums []int, path []int, check map[int]bool, res *[][]int) {
	if len(path) == len(nums) {
		if !isExist(res, path) {
			*res = append(*res, append([]int{}, path...))
		}
		return
	}
	for i, _ := range nums {
		if val, _ := check[i]; !val {
			check[i] = true
			dfs(nums, append(path, nums[i]), check, res)
			check[i] = false
		}
	}
}

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	check := map[int]bool{}
	for i, _ := range nums {
		check[i] = false
	}
	dfs(nums, []int{}, check, &res)
	return res
}
