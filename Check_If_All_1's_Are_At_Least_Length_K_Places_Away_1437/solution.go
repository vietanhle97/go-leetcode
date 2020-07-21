package Check_If_All_1_s_Are_At_Least_Length_K_Places_Away_1437

func kLengthApart(nums []int, k int) bool {

	prev := -1

	for i, e := range nums {
		if e == 1 && prev == -1 {
			prev = i
			continue
		} else if e == 1 {
			if i-prev <= k {
				return false
			}
			prev = i
		}
	}
	return true
}
