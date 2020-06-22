package Sum_Of_Mutated_Array_Closest_To_Target_1300

import "sort"

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	curSum := 0
	m := len(arr)
	for i, val := range arr {
		tmp := curSum + (m-i)*val
		if tmp == target {
			return val
		}
		if tmp < target {
			curSum += val
		}
		if tmp > target {
			v := (target - curSum) / (m - i)
			if abs(target-curSum+v*(m-i)) <= abs(target-curSum+(v+1)*(m-i)) {
				return v
			}
			return v + 1
		}
	}
	return arr[m-1]
}
