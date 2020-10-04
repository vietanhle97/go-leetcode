package Frog_Jump_403

func bs(l, r, v, k int, arr []int) int {
	for l <= r {
		mid := l + (r-l)/2
		if v-arr[mid] == k {
			return mid
		} else if v-arr[mid] < k {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
func canCross(stones []int) bool {
	if stones[1] > 1 {
		return false
	}
	m := len(stones)
	if m == 2 {
		return true
	}
	if stones[2]-stones[1] > 2 {
		return false
	}
	dp := make([][]bool, m)
	for i := range dp {
		dp[i] = make([]bool, m)
	}
	dp[1][0] = true
	dp[2][1] = true
	for i := 3; i < m; i++ {
		for j := 1; j < i; j++ {
			k := stones[i] - stones[j]
			a, b, c := bs(1, j, stones[j], k-1, stones), bs(1, j, stones[j], k, stones), bs(1, j, stones[j], k+1, stones)
			if a != -1 && dp[j][a] {
				dp[i][j] = dp[j][a]
			}
			if b != -1 && dp[j][b] {
				dp[i][j] = dp[j][b]
			}
			if c != -1 && dp[j][c] {
				dp[i][j] = dp[j][c]
			}
			if i == m-1 && dp[i][j] {
				return true
			}
		}
	}
	return false
}
