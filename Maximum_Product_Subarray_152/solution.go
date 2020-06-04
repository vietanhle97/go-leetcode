package Maximum_Product_Subarray_152

func max(n []int) int {
	res := n[0]
	for i := range n {
		if n[i] > res {
			res = n[i]
		}
	}
	return res
}

func min(n []int) int {
	res := n[0]
	for i := range n {
		if n[i] < res {
			res = n[i]
		}
	}
	return res
}
func maxProduct(nums []int) int {
	a, b, c := nums[0], nums[0], nums[0]
	for _, i := range nums[1:] {
		b, c = max([]int{i, i * b, i * c}), min([]int{i, i * b, i * c})
		a = max([]int{a, b})
	}
	return a
}
