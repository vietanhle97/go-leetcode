package Minimum_Operations_To_Reduce_X_To_Zero_1658

func minOperations(nums []int, x int) int {
	rest := 0
	for _, e := range nums {
		rest += e
	}
	rest -= x
	tmp, st, en, res := 0, 0, 0, make([]int, 0)
	for _, e := range nums {
		tmp += e
		if tmp <= rest {
			en++
			if tmp == rest {
				res = append(res, en-st)
			}
		} else {
			en++
			for tmp > rest && st < en {
				tmp -= nums[st]
				st++
				if tmp == rest {
					res = append(res, en-st)
				}
			}
		}
	}
	if len(res) == 0 {
		return -1
	}
	ans := 0
	for _, e := range res {
		if e > ans {
			ans = e
		}
	}
	return len(nums) - ans
}
