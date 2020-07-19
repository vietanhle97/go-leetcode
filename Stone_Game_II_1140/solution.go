package Stone_Game_II_1140

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

func dfs(start int, M int, sum []int, piles []int, vis *[100][100]int) int {
	if start >= len(piles) {
		return 0
	}
	if (*vis)[start][M] > 0 {
		return (*vis)[start][M]
	}

	total := sum[start]
	res := 0

	for X := 1; X < min(2*M+1, len(piles)-start+1); X++ {
		op := dfs(start+X, max(M, X), sum, piles, vis)
		res = max(total-op, res)
	}
	(*vis)[start][M] = res
	return res
}

func stoneGameII(piles []int) int {

	n := len(piles)
	vis := [100][100]int{}
	sum := make([]int, n)
	sum[n-1] = piles[n-1]
	for i := n - 2; i >= 0; i-- {
		sum[i] = piles[i] + sum[i+1]
	}

	dfs(0, 1, sum, piles, &vis)
	return vis[0][1]
}
