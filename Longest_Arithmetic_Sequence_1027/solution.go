package Longest_Arithmetic_Sequence_1027

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func longestArithSeqLength(A []int) int {
	res := 0
	dp := make([][]int, len(A))
	max_ := A[0]
	for _, e := range A {
		if abs(e) > abs(max_) {
			max_ = abs(e)
		}
	}
	for i, _ := range dp {
		dp[i] = make([]int, max_+max_+1)
	}

	for i, v1 := range A {
		for j, v2 := range A[:i] {
			d := v2 - v1 + max_
			if dp[j][d] != 0 {
				dp[i][d] = dp[j][d] + 1
			} else {
				dp[i][d] = 2
			}
			if dp[i][d] > res {
				res = dp[i][d]
			}
		}
	}
	return res
}
