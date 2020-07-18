package The_Kth_Factor_Of_N_1492_

import "sort"

func kthFactor(n int, k int) int {
	fac := make([]int, 0)
	d := 1
	for d*d <= n {
		if n%d == 0 {
			fac = append(fac, d)
			if d*d != n {
				fac = append(fac, n/d)
			}
		}
		d++
	}
	// fmt.Println(fac)
	sort.Ints(fac)
	if len(fac) < k {
		return -1
	}
	return fac[k-1]
}
