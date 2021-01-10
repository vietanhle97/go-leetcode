package Jump_Game_VI_1696

func maxResult(nums []int, k int) int {
	m := len(nums)
	dp, q := make([]int, m), [][]int{[]int{nums[0], 0}}
	dp[0] = nums[0]
	for i := 1; i < m; i++ {
		dp[i] = nums[i] + q[0][0]
		for len(q) > 0 && q[len(q)-1][0] < dp[i] {
			q = q[:len(q)-1]
		}
		q = append(q, []int{dp[i], i})
		if i-k == q[0][1] {
			q = q[1:]
		}
	}
	return dp[m-1]
}
