package Shortest_Subarray_With_Sum_At_Least_K_862

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func shortestSubarray(A []int, K int) int {
	pref, N := []int{0}, len(A)
	for i := range A {
		pref = append(pref, A[i]+pref[len(pref)-1])
	}
	q, ans := make([]int, 0), N+1

	for i, e := range pref {

		for len(q) > 0 && e <= pref[q[len(q)-1]] {
			q = q[:len(q)-1]
		}

		for len(q) > 0 && e-pref[q[0]] >= K {
			ans = min(ans, i-q[0])
			q = q[1:]
		}
		q = append(q, i)
	}
	if ans < N+1 {
		return ans
	}
	return -1
}
