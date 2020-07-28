package Subarray_Sums_Divisible_By_K_974

func subarraysDivByK(A []int, k int) int {
	m := len(A)
	sum := make([]int, m+1)
	div := make([]int, k)
	res := 0
	for i := 0; i < m+1; i++ {
		if i > 0 {
			sum[i] = sum[i-1] + A[i-1]
		}
		if sum[i]%k >= 0 {
			div[sum[i]%k]++
		} else {
			div[sum[i]%k+k]++
		}

	}
	for _, e := range div {
		if e > 1 {
			res += e * (e - 1) / 2
		}
	}
	return res
}
