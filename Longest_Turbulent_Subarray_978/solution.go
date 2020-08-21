package Longest_Turbulent_Subarray_978

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxTurbulenceSize(A []int) int {
	m, ans := len(A), 1
	even, odd := make([]int, m), make([]int, m)
	even[0], odd[0] = 1, 1
	for i := 1; i < m; i++ {
		if i%2 == 0 {
			if A[i-1] > A[i] {
				odd[i], even[i] = odd[i-1]+1, 1
			} else if A[i-1] < A[i] {
				odd[i], even[i] = 1, even[i-1]+1
			} else {
				odd[i], even[i] = 1, 1
			}
		} else {
			if A[i] > A[i-1] {
				odd[i], even[i] = odd[i-1]+1, 1
			} else if A[i] < A[i-1] {
				odd[i], even[i] = 1, even[i-1]+1
			} else {
				odd[i], even[i] = 1, 1
			}
		}
		ans = max(ans, max(even[i], odd[i]))
	}
	return ans
}
