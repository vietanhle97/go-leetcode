package _32_Pattern_456

func find132pattern(nums []int) bool {
	max := -1 << 31
	stack := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < max {
			return true
		}
		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			max = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}
