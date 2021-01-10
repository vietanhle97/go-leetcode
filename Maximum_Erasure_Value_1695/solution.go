package Maximum_Erasure_Value_1695

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumUniqueSubarray(nums []int) int {
	m := len(nums)
	pref := make([]int, m)
	for i := range nums {
		if i == 0 {
			pref[i] = nums[i]
		} else {
			pref[i] = nums[i] + pref[i-1]
		}
	}
	start, end, res, check := 0, 0, 0, map[int]int{}
	for i := range nums {
		if _, ok := check[nums[i]]; !ok {
			check[nums[i]] = i
			end = i
			if start == 0 {
				res = max(res, pref[end])
			} else {
				res = max(res, pref[end]-pref[start-1])
			}

		} else {
			start = max(start, check[nums[i]]+1)
			end = i
			check[nums[i]] = i
			if start == 0 {
				res = max(res, pref[end])
			} else {
				res = max(res, pref[end]-pref[start-1])
			}
		}
	}
	return res
}
