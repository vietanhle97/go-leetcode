package Max_Dot_Product_Of_Two_Subsequences_1458

func min(A []int) int {
	res := A[0]
	for _, e := range A {
		if e < res {
			res = e
		}
	}
	return res
}

func max(A []int) int {
	res := A[0]
	for _, e := range A {
		if e > res {
			res = e
		}
	}
	return res
}

func allNeg(A []int) bool {
	for _, e := range A {
		if e > 0 {
			return false
		}
	}
	return true
}

func allPos(A []int) bool {
	for _, e := range A {
		if e < 0 {
			return false
		}
	}
	return true
}

func maxDotProduct(A []int, B []int) int {
	if allNeg(A) && allPos(B) {
		return max(A) * min(B)
	}

	if allPos(A) && allNeg(B) {
		return max(B) * min(A)
	}

	m, n := len(A), len(B)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		if i == 0 {
			continue
		}
		for j := 1; j < n+1; j++ {
			dp[i][j] = max([]int{dp[i][j-1], dp[i-1][j], dp[i-1][j-1] + A[i-1]*B[j-1]})
		}
	}
	return dp[m][n]
}
