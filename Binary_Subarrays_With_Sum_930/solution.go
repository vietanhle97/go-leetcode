package Binary_Subarrays_With_Sum_930

func numSubarraysWithSum(A []int, S int) int {
	P := []int{0}
	sum := 0
	cnt := 0
	for i := range A {
		sum += A[i]
		P = append(P, sum)
	}
	check := make([]int, sum+S+1)
	for i := range P {
		cnt += check[P[i]]
		check[P[i]+S] += 1
	}
	return cnt
}
