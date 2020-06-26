package Smallest_Integer_Divisible_By_K_1015

func smallestRepunitDivByK(K int) int {
	if K%2 == 0 || K%5 == 0 {
		return -1
	}
	N := 1
	cnt := 1
	for N%K != 0 {
		N = (N%K)*10 + 1
		cnt++
	}
	return cnt
}
