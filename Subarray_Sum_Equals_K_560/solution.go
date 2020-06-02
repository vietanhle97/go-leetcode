package Subarray_Sum_Equals_K_560

func subarraySum(nums []int, k int) int {
	check := map[int]int{}
	check[0] = 1
	s := 0
	cnt := 0
	for i, _ := range nums {
		s += nums[i]
		if _, ok := check[s-k]; ok {
			cnt += check[s-k]
		}
		if _, ok := check[s]; ok {
			check[s] += 1
		} else {
			check[s] = 1
		}
	}
	return cnt
}
