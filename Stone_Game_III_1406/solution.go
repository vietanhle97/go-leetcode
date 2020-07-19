package Stone_Game_III_1406

const (
	MaxInt = 1<<31 - 1
	MinInt = -MaxInt
)

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

func dfs(start int, piles []int, vis *[]int) int {
	if start >= len(piles) {
		return 0
	}
	if (*vis)[start] != MaxInt {
		return (*vis)[start]
	}

	sum := 0
	res := MinInt

	for i := start; i < min(start+3, len(piles)); i++ {
		sum += piles[i]
		res = max(res, sum-dfs(i+1, piles, vis))
	}
	(*vis)[start] = res
	return res
}

func stoneGameIII(piles []int) string {
	n := len(piles)
	vis := make([]int, n)
	for i := range vis {
		vis[i] = MaxInt
	}
	dfs(0, piles, &vis)
	if vis[0] < 0 {
		return "Bob"
	} else if vis[0] > 0 {
		return "Alice"
	}
	return "Tie"
}
