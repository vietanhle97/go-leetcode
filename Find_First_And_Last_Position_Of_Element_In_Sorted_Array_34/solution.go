package Find_First_And_Last_Position_Of_Element_In_Sorted_Array_34

func binarySearch(nums []int, target int, ls bool) int {
	m := len(nums)
	lo, hi := 0, m
	for lo < hi {
		mid := lo + (hi-lo)/2
		if (nums[mid] == target && ls) || nums[mid] > target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	start := binarySearch(nums, target, true)
	if start == len(nums) || nums[start] != target {
		return res
	}
	end := binarySearch(nums, target, false) - 1
	res[0], res[1] = start, end
	return res
}
