package Minimum_Cost_For_Tickets_983

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a ...int) int {
	res := a[0]
	for _, e := range a {
		if e < res {
			res = e
		}
	}
	return res
}

func mincostTickets(days []int, costs []int) int {
	m := map[int]bool{}
	for _, e := range days {
		m[e] = true
	}
	n := days[len(days)-1]
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if !m[i] {
			dp[i] = dp[i-1]
			continue
		}
		dp[i] = min(dp[i-1]+costs[0], dp[max(i-7, 0)]+costs[1], dp[max(i-30, 0)]+costs[2])

	}
	return dp[n]
}
