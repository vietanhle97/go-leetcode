package Allocate_Mailboxes_1478

import "sort"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dist(i, j int, B []int) int {
	m1, m2 := (i+j)/2, (i+j+1)/2
	return B[j+1] - B[m2] - (B[m1+1] - B[i])
}

func minDistance(houses []int, K int) int {
	sort.Ints(houses)
	N := len(houses)
	sum := 0
	D := make([]int, N+1)
	for i := range houses {
		sum += houses[i]
		D[i+1] = sum
	}

	dp := make([]int, N)
	for j := 0; j < N; j++ {
		dp[j] = dist(0, j, D)
	}

	for i := 2; i <= K; i++ {
		for j := N - 1; j >= i-1; j-- {
			for k := i - 2; k < j; k++ {
				dp[j] = min(dp[j], dp[k]+dist(k+1, j, D))
			}
		}
	}

	return dp[N-1]
}
