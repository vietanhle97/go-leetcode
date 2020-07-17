package Single_Element_In_A_Sorted_Array_540

func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := 2 * ((l + r) / 4)
		if nums[m] == nums[m+1] {
			l = m + 2
		} else {
			r = m
		}
	}
	return nums[l]
}
