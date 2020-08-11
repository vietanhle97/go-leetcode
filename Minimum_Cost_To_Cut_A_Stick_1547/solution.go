package Minimum_Cost_To_Cut_A_Stick_1547

import "sort"

const MaxInt = 1<<31 - 1

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCost(n int, cuts []int) int {
	sort.Ints(cuts)
	m := len(cuts)
	dp := make([][]int, m+2)
	cuts = append([]int{0}, cuts...)
	cuts = append(cuts, n)
	for i := range dp {
		dp[i] = make([]int, m+2)
		for j := range dp[i] {
			dp[i][j] = MaxInt
		}
	}

	for i := 0; i <= m+1; i++ {
		dp[i][i] = 0
	}
	for i := m; i >= 1; i-- {
		for j := i + 1; j <= m+1; j++ {
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j]+cuts[j]-cuts[i-1])
			}
		}

	}
	return dp[1][m+1]
}
