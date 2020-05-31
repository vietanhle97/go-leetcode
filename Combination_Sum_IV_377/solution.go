package Combination_Sum_IV_377

func findMin(nums []int) int {
	min_ := nums[0]
	for _, e := range nums {
		if e < min_ {
			min_ = e
		}
	}
	return min_
}

func combinationSum4(nums []int, target int) int {
	m := len(nums)
	if m < 1 {
		return 0
	}
	min_ := findMin(nums)
	if min_ > target {
		return 0
	}
	table := make([]int, target+1)
	table[min_] = 1
	for i := 0; i < min_; i++ {
		table[i] = -1
	}
	for i := min_ + 1; i < target+1; i++ {
		for _, e := range nums {
			if e < i {
				if table[i-e] != -1 {
					table[i] += table[i-e]
				}
			} else if e == i {
				table[i] += 1
			}
		}
		if table[i] == 0 {
			table[i] = -1
		}
	}
	if table[target] == -1 {
		return 0
	}
	return table[target]
}
