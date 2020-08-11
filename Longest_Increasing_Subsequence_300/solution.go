package Longest_Increasing_Subsequence_300

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
func lengthOfLIS(nums []int) int {
	m := len(nums)
	if m <= 1 {
		return m
	}
	ans := 0
	table := make([]int, m)
	table[0] = 1
	for i := 1; i < m; i++ {
		max_ := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				max_ = max(table[j], max_)
			}
		}
		table[i] = max_ + 1
		ans = max(ans, table[i])
	}
	return ans
}
