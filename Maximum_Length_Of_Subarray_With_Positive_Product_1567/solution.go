package Maximum_Length_Of_Subarray_With_Positive_Product_1567

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func getMaxLen(nums []int) int {
	odd, even, cnt, res := -1, -1, 0, 0
	for i := range nums {
		if nums[i] == 0 {
			even, odd, cnt = i, -1, 0
		}
		if nums[i] < 0 {
			cnt++
			if odd == -1 {
				odd = i
			}
		}
		if cnt%2 == 1 {
			res = max(res, i-odd)
		} else {
			res = max(res, i-even)
		}
	}
	return res
}
