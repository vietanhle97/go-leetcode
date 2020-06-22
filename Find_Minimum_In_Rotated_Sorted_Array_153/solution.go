package Find_Minimum_In_Rotated_Sorted_Array_153

func findRotatedPivot(nums []int, l, r int) int {
	if l > r {
		return -1
	}
	mid := l + (r-l)/2
	if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
		return mid + 1
	}
	if mid > 0 && nums[mid] < nums[mid-1] {
		return mid
	}
	left := findRotatedPivot(nums, l, mid-1)
	if left != -1 {
		return left
	}
	return findRotatedPivot(nums, mid+1, r)
}

func findMin(nums []int) int {
	pivot := findRotatedPivot(nums, 0, len(nums)-1)
	if pivot == -1 {
		return nums[0]
	}
	return nums[pivot]
}
