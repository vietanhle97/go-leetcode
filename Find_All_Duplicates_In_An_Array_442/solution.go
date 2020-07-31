package Find_All_Duplicates_In_An_Array_442

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	for _, e := range nums {
		// fmt.Println(nums)
		if nums[abs(e)-1] < 0 {
			res = append(res, abs(e))
		} else {
			nums[abs(e)-1] = -nums[abs(e)-1]
		}
	}
	return res
}
