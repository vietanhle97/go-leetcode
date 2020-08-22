package Minimum_Numbers_Of_Function_Calls_To_Make_Target_Array_1558

func minOperations(target []int) int {
	res := 0
	n := len(target)
	for {
		cnt, i := 0, 0
		for i < n {
			if (target[i] & 1) > 0 {
				break
			} else if target[i] == 0 {
				cnt++
			}
			i++

			if cnt == n {
				return res
			}
			if i == n {
				for j := 0; j < n; j++ {
					target[j] = target[j] / 2
				}
				res++
			}
		}
		for j := i; j < n; j++ {
			if (target[j] & 1) != 0 {
				target[j]--
				res++
			}

		}
	}
	return 0
}
