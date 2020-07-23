package Jump_Game_II_45

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func jump(nums []int) int {
	m := len(nums)
	dp := make([]int, m+1)
	dp[1] = 0

	for j := 1; j < m+1; j++ {
		if dp[m] == 1 {
			return 1
		}
		for k := 1; k <= nums[j-1]; k++ {
			if j+k < m+1 && dp[j+k] == 0 {
				dp[j+k] = dp[j] + 1
			} else if j+k < m+1 {
				dp[j+k] = min(dp[j+k], dp[j]+1)
			}
		}
	}
	// fmt.Println(dp)
	return dp[m]
}
