package Video_Stitching_1024

import "sort"

const MaxInt = 1<<31 - 1

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func videoStitchingSol2(clips [][]int, T int) int {
	sort.Slice(clips, func(i, j int) bool {
		if clips[i][0] == clips[j][0] {
			return clips[i][0] < clips[j][0]
		}
		return clips[i][0] < clips[j][0]
	})

	dp := make([]int, T+1)

	for i := range dp {
		dp[i] = MaxInt
	}
	dp[0] = 0

	for i := 1; i < T+1; i++ {
		for _, e := range clips {
			if e[1] >= i && i > e[0] {
				dp[i] = min(dp[i], dp[e[0]]+1)
			}
		}
	}
	if dp[T] == MaxInt {
		return -1
	}
	return dp[T]
}
