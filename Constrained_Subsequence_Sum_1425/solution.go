package Constrained_Subsequence_Sum_1425

type p struct {
	ind int
	val int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func constrainedSubsetSum(nums []int, k int) int {
	m := len(nums)
	res, dp := nums[0], []p{p{0, nums[0]}}
	for i := 1; i < m; i++ {
		tmp := nums[i] + max(0, dp[0].val)
		if dp[0].ind+k == i {
			dp = dp[1:]
		}
		for len(dp) > 0 && tmp >= dp[len(dp)-1].val {
			dp = dp[:len(dp)-1]
		}
		dp = append(dp, p{i, tmp})
		res = max(res, tmp)
	}
	return res
}
