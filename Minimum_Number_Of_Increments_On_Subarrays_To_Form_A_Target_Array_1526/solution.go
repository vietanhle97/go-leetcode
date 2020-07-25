package Minimum_Number_Of_Increments_On_Subarrays_To_Form_A_Target_Array_1526

func minNumberOperations(target []int) int {
	n := len(target)
	res := make([]int, n)
	res[0] = target[0]
	cnt := target[0]
	for i := 1; i < n; i++ {
		if target[i] >= target[i-1] {
			tmp := res[i-1]
			if tmp == 0 {
				res[i] = target[i] - target[i-1]
				cnt += res[i]
			} else {
				cnt -= res[i-1]
				res[i-1] = 0
				res[i] = target[i] - target[i-1] + tmp
				cnt += res[i]
			}

		} else {
			res[i] = 0
		}
	}
	return cnt
}
