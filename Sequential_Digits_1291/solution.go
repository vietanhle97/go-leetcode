package Sequential_Digits_1291

import "sort"

func sequentialDigits(low int, high int) []int {
	first := low
	cnt := 0
	for first >= 10 {
		first /= 10
		cnt++
	}
	cnt++
	res := make([]int, 0)
	for i := 1; i <= 10-cnt; i++ {
		cur, tmp := i, i+1
		for tmp <= 9 {
			cur = 10*cur + tmp
			tmp++
			if cur > high {
				break
			}
			if cur >= low && cur <= high {
				res = append(res, cur)
			}
		}
	}
	sort.Ints(res)
	return res
}
