package Sum_Of_Absolute_Differences_In_A_Sorted_Array_1685

func getSumAbsoluteDifferences(nums []int) []int {
	sum, m, pref := 0, len(nums), make([]int, len(nums))
	for i, e := range nums {
		if i == 0 {
			pref[i] = e
		} else {
			pref[i] = pref[i-1] + e
		}
		sum += e
	}
	res := make([]int, m)
	for i, e := range nums {
		res[i] = e*(i+1) - pref[i] + sum - pref[i] - e*(m-i-1)
	}
	return res
}
