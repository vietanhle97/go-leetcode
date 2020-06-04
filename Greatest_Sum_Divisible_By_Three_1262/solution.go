package Greatest_Sum_Divisible_By_Three_1262

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxSumDivThree(nums []int) int {
	seen := []int{0, 0, 0}
	for i := range nums {
		tmp := make([]int, len(seen))
		copy(tmp, seen)
		for e := range tmp {
			seen[(nums[i]+tmp[e])%3] = max(seen[(nums[i]+tmp[e])%3], nums[i]+tmp[e])
		}
	}
	return seen[0]
}
