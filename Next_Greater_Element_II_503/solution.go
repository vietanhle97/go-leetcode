package Next_Greater_Element_II_503

func nextGreaterElements(nums []int) []int {
	m := len(nums)
	if m == 0 {
		return []int{}
	}
	if m < 2 {
		return []int{-1}
	}
	nums = append(nums, nums...)
	stack := make([]int, 0)
	res := make([]int, 2*m)
	for i := range nums {
		if i == 2*m-1 {
			res[i] = -1
		}
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
				res[stack[len(stack)-1]] = nums[i]
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
		}
	}
	if len(stack) > 0 {
		for i := range stack {
			res[stack[i]] = -1
		}
	}
	return res[:m]
}
