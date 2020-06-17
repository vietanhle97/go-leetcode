package Product_Of_Array_Except_Self_238

func productExceptSelf(nums []int) []int {
	m := len(nums)
	ans := make([]int, m)
	b := 1
	for i, e := range nums {
		ans[i] = b
		b *= e
	}
	b = 1
	for i := m - 1; i > -1; i-- {
		ans[i] = ans[i] * b
		b *= nums[i]
	}
	return ans
}
