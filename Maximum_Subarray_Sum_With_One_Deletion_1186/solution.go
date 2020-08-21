package Maximum_Subarray_Sum_With_One_Deletion_1186

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumSum(A []int) int {
	n := len(A)
	ans, noDel, del := A[0], A[0], 0

	for i := 1; i < n; i++ {
		del = max(del+A[i], noDel)
		noDel = max(noDel+A[i], A[i])
		ans = max(ans, max(noDel, del))
	}
	return ans
}
