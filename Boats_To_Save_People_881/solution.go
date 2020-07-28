package Boats_To_Save_People_881

import "sort"

func numRescueBoats(ppl []int, limit int) int {
	sort.Ints(ppl)
	i, j := 0, len(ppl)-1
	tmp, res := limit, 0
	cnt := 2
	for i <= j {
		for ppl[j] <= tmp && cnt > 0 {
			tmp -= ppl[j]
			j--
			cnt--
			if j < 0 || j < i {
				return res + 1
			}
		}
		if cnt == 0 || tmp == 0 {
			tmp, cnt = limit, 2
			res++
		} else {
			if tmp >= ppl[i] {
				tmp -= ppl[i]
				i++
				if i > j || i > len(ppl)-1 {
					return res + 1
				}
			}
			cnt, tmp = 2, limit
			res++
		}
	}
	return res
}
