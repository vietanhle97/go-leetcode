package Sliding_Window_Maximum_239

func maxSlidingWindowSol2(nums []int, k int) []int {
	n, q, res := len(nums), make([]int, 0), make([]int, 0)
	for i := 0; i < k; i++ {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	for i := k; i < n; i++ {
		res = append(res, nums[q[0]])
		for len(q) > 0 && q[0] <= i-k {
			q = q[1:]
		}
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	res = append(res, nums[q[0]])
	return res
}
