package Find_Peak_Element_162

func findPeakElement(nums []int) int {
	lo, hi := 0, len(nums)-1
	if len(nums) == 1 {
		return 0
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return 0
		}
		return 1
	}
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if mid == 0 || mid == len(nums)-1 {
			return mid
		}
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		} else {
			if nums[mid] < nums[mid+1] {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
	}
	return nums[lo]
}
