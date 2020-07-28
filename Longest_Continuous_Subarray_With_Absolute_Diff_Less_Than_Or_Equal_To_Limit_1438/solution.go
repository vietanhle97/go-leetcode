package Longest_Continuous_Subarray_With_Absolute_Diff_Less_Than_Or_Equal_To_Limit_1438

func longestSubarray(nums []int, limit int) int {
	maxd, mind := make([]int, 0), make([]int, 0)
	n := 0
	for i := 0; i < len(nums); i++ {
		for len(maxd) != 0 && nums[i] > maxd[len(maxd)-1] {
			maxd = maxd[:len(maxd)-1]
		}
		for len(mind) != 0 && nums[i] < mind[len(mind)-1] {
			mind = mind[:len(mind)-1]
		}
		maxd = append(maxd, nums[i])
		mind = append(mind, nums[i])
		if maxd[0]-mind[0] > limit {
			if maxd[0] == nums[n] {
				maxd = maxd[1:]
			}
			if mind[0] == nums[n] {
				mind = mind[1:]
			}
			n++
		}
	}
	return len(nums) - n
}
