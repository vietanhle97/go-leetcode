package Number_Of_Longest_Increasing_Subsequence_673

func findNumberOfLIS(nums []int) int {
	m := len(nums)
	dp, cnt := make([]int, m), make([]int, m)
	max := 0
	for i := range cnt {
		cnt[i] = 1
	}
	for i := range nums {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j] >= dp[i] {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[j]+1 == dp[i] {
					cnt[i] += cnt[j]
				}
			}
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	ans := 0
	for i := range dp {
		if dp[i] == max {
			ans += cnt[i]
		}
	}
	return ans
}
