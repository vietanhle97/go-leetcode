package Sum_Of_Subsequence_Widths_891

import "sort"

const MOD = int(1e9) + 7

func sumSubseqWidths(A []int) int {
	m := len(A)
	sort.Ints(A)

	pow := []int{1}
	for i := 1; i <= m; i++ {
		pow = append(pow, (pow[len(pow)-1]*2)%MOD)
	}
	res := 0
	for i, e := range A {
		res = (res + e*(pow[i]-pow[m-i-1])) % MOD
	}
	return res
}
