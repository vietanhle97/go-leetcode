package House_Robber_II_213

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func rob(nums []int) int {
	m := len(nums)
	if m == 0 {
		return 0
	}
	if m == 1 {
		return nums[0]
	}
	dp1 := make([]int, m)
	dp2 := make([]int, m)
	res := 0
	for i := 0; i < m-1; i++ {
		dp1[i] = nums[i]
		for j := i - 2; j >= 0; j-- {
			dp1[i] = max(dp1[i], dp1[j]+nums[i])
		}
		res = max(dp1[i], res)
	}
	for i := 1; i < m; i++ {
		dp2[i] = nums[i]
		for j := i - 2; j >= 0; j-- {
			dp2[i] = max(dp2[i], dp2[j]+nums[i])
		}
		res = max(dp2[i], res)
	}
	return res
}
