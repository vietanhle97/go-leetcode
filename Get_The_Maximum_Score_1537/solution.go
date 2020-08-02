package Get_The_Maximum_Score_1537

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxSum(A []int, B []int) int {
	m, n := len(A), len(B)
	a, b := 0, 0
	MOD := int(1e9) + 7
	res := 0
	i, j := 0, 0
	for i < m || j < n {
		if i < m && (j >= n || A[i] < B[j]) {
			a += A[i]
			i++
		} else if j < n && (i >= m || A[i] > B[j]) {
			b += B[j]
			j++

		} else {
			res += max(a, b) + A[i]
			i++
			j++
			a, b = 0, 0
		}
	}
	return (res + max(a, b)) % MOD
}
