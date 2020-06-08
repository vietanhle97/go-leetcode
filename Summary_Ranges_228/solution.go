package Summary_Ranges_228

import "strconv"

func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	if len(nums) == 0 {
		return res
	}
	start := nums[0]
	prev := nums[0]
	end := false
	if len(nums) == 1 {
		res = append(res, strconv.Itoa(start))
		return res
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] == prev+1 {
			prev = nums[i]
			if i == len(nums)-1 {
				res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(prev))
				end = true
			}
			continue
		} else {
			if start != prev {
				res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(prev))
			} else {
				res = append(res, strconv.Itoa(start))
			}

			start = nums[i]
			prev = nums[i]
		}
	}
	if !end {
		res = append(res, strconv.Itoa(start))
	}
	return res
}
