package Minimum_Deletion_Cost_To_Avoid_Repeating_Letters_1578

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCost(s string, cost []int) int {
	prev, ans := -1, 0
	for i := range s {
		if prev == -1 {
			prev = i
		} else {
			if s[i] == s[prev] {
				if cost[i] > cost[prev] {
					ans += cost[prev]
					prev = i
				} else {
					ans += cost[i]
				}
			} else {
				prev = i
			}
		}
	}
	return ans
}
