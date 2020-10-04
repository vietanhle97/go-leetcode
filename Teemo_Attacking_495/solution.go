package Teemo_Attacking_495

func findPoisonedDuration(T []int, d int) int {
	st := make([]int, 0)
	ans := 0
	for _, e := range T {
		if len(st) == 0 {
			st = append(st, e+d-1)
			ans += d
		} else {
			if e <= st[len(st)-1] {
				ans += e + d - 1 - st[len(st)-1]
				st[len(st)-1] = e + d - 1
			} else {
				st = append(st, e+d-1)
				ans += d
			}
		}
	}
	return ans
}
