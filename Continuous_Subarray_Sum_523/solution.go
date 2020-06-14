package Continuous_Subarray_Sum_523

func checkSubarraySum(nums []int, k int) bool {
	seen, cur := map[int]int{0: -1}, 0
	for i, e := range nums {
		if k == 0 {
			cur += e
		} else {
			cur = (cur + e) % k
		}
		if _, ok := seen[cur]; ok {
			if i-seen[cur] > 1 {
				return true
			}
		} else {
			seen[cur] = i
		}
	}
	return false
}
