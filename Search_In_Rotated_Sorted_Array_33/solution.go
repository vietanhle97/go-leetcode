package Search_In_Rotated_Sorted_Array_33

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

func binarySearch(nums []int, l, r, target int) int {
	if l > r {
		return -1
	}
	mid := l + (r-l)/2
	if nums[mid] == target {
		return mid
	}
	left := binarySearch(nums, l, mid-1, target)
	if left != -1 {
		return left
	}
	return binarySearch(nums, mid+1, r, target)
}

func search(nums []int, target int) int {
	pivot := findRotatedPivot(nums, 0, len(nums)-1)
	if pivot == -1 {
		return binarySearch(nums, 0, len(nums)-1, target)
	} else {
		ls := binarySearch(nums, 0, pivot-1, target)
		if ls != -1 {
			return ls
		}
		return binarySearch(nums, pivot, len(nums)-1, target)
	}
}
