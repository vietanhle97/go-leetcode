package Minimum_Operations_To_Make_A_Subsequence_1713

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func bs(A []int, l, r, v int) int {
	for r-l > 1 {
		mid := l + (r-l)/2
		if A[mid] >= v {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

func minOperations(target []int, arr []int) int {
	m := map[int]int{}
	for i, e := range target {
		m[e] = i
	}
	A := make([]int, 0)
	for _, e := range arr {
		if _, ok := m[e]; ok {
			A = append(A, m[e])
		}
	}
	l, cnt := len(A), 0
	dp := make([]int, l+1)
	dp[0] = A[0]
	cnt++

	for i := 1; i < l; i++ {
		if A[i] < dp[0] {
			dp[0] = A[i]
		} else if A[i] > dp[cnt-1] {
			dp[cnt] = A[i]
			cnt++
		} else {
			dp[bs(dp, -1, cnt-1, A[i])] = A[i]
		}
	}
	return len(target) - cnt
}
