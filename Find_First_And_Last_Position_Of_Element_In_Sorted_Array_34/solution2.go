package Find_First_And_Last_Position_Of_Element_In_Sorted_Array_34

func bsLeft(l, r, target int, nums []int) int {
	if l > r {
		return -1
	}
	if nums[l] == target {
		return l
	}
	mid := l + (r-l)/2
	ls := bsLeft(l, mid-1, target, nums)
	if ls != -1 {
		return ls
	}
	if nums[mid] == target {
		return mid
	}
	return bsLeft(mid+1, r, target, nums)
}

func bsRight(l, r, target int, nums []int) int {
	if l > r {
		return -1
	}
	if nums[r] == target {
		return r
	}
	mid := l + (r-l)/2
	rs := bsRight(mid+1, r, target, nums)
	if rs != -1 {
		return rs
	}
	if nums[mid] == target {
		return mid
	}
	return bsRight(l, mid-1, target, nums)
}

func searchRangeSol2(nums []int, target int) []int {
	m := len(nums)
	if m == 0 {
		return []int{-1, -1}
	}
	l, r := bsLeft(0, m-1, target, nums), bsRight(0, m-1, target, nums)
	return []int{l, r}
}
