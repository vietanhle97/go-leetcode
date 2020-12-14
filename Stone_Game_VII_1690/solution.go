package Stone_Game_VII_1690

var pref []int
var dp [][]int

const MaxInt = 1<<31 - 1

func sum(l, r int) int {
	return pref[r+1] - pref[l]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func DP(l, r, isAlice int) int {
	if l > r {
		return 0
	}
	if dp[l][r] != 0 {
		return dp[l][r]
	}

	if isAlice == 0 {
		a := DP(l+1, r, 1-isAlice) + sum(l+1, r)
		b := DP(l, r-1, 1-isAlice) + sum(l, r-1)
		dp[l][r] = max(a, b)
	} else {
		a := DP(l+1, r, 1-isAlice) - sum(l+1, r)
		b := DP(l, r-1, 1-isAlice) - sum(l, r-1)
		dp[l][r] = min(a, b)
	}
	return dp[l][r]
}

func stoneGameVII(stones []int) int {
	m := len(stones)
	pref = make([]int, m+1)
	for i := 0; i < m; i++ {
		pref[i+1] = pref[i] + stones[i]
	}
	dp = make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	return DP(0, m-1, 0)
}
