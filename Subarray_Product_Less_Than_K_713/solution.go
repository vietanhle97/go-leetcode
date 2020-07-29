package Subarray_Product_Less_Than_K_713

func numSubarrayProductLessThanK(nums []int, k int) int {
	m := len(nums)
	prod := make([]int, m+1)
	sum := 1
	res := 0
	for i := 1; i < m+1; i++ {
		if nums[i-1] >= k {
			sum = 1
			continue
		}
		sum *= nums[i-1]
		if sum < k {
			prod[i] = prod[i-1] + 1
		} else {
			cnt := prod[i-1] + 1
			j := i - cnt
			for sum >= k {
				sum /= nums[j]
				cnt--
				j++
			}
			prod[i] = cnt
		}
		res += prod[i]
	}
	return res
}
