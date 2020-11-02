package Best_Team_With_No_Conflicts_1626

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func bestTeamScore(scores []int, ages []int) int {
	l := len(scores)
	m := make([][]int, 0)
	for i := range scores {
		m = append(m, []int{ages[i], scores[i]})
	}
	sort.Slice(m, func(i, j int) bool {
		if m[i][0] == m[j][0] {
			return m[i][1] < m[j][1]
		}
		return m[i][0] < m[j][0]
	})

	res := 0
	dp := make([]int, l)
	dp[0] = m[0][1]
	res = max(res, dp[0])
	for i := 1; i < l; i++ {
		dp[i] = m[i][1]
		for j := 0; j < i; j++ {
			if m[j][1] <= m[i][1] {
				dp[i] = max(dp[i], dp[j]+m[i][1])
			}
		}
		res = max(res, dp[i])
	}
	return res
}
