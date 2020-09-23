package Maximum_Non_Negative_Product_In_A_Matrix_1594

const mod = int(1e9) + 7

func min(a ...int) int {
	res := a[0]
	for _, e := range a {
		if e < res {
			res = e
		}
	}
	return res
}

func max(a ...int) int {
	res := a[0]
	for _, e := range a {
		if e > res {
			res = e
		}
	}
	return res
}

func maxProductPath(grid [][]int) int {
	r := len(grid)
	c := len(grid[0])
	dp := make([][][]int, r)

	for i := range dp {
		dp[i] = make([][]int, c)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}

	dp[0][0][0] = grid[0][0]
	dp[0][0][1] = grid[0][0]

	for i := 1; i < c; i++ {
		dp[0][i][0] = min(dp[0][i-1][0]*grid[0][i], dp[0][i-1][1]*grid[0][i])
		dp[0][i][1] = max(dp[0][i-1][0]*grid[0][i], dp[0][i-1][1]*grid[0][i])
	}

	for i := 1; i < r; i++ {
		dp[i][0][0] = min(dp[i-1][0][0]*grid[i][0], dp[i-1][0][1]*grid[i][0])
		dp[i][0][1] = max(dp[i-1][0][0]*grid[i][0], dp[i-1][0][1]*grid[i][0])
	}

	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			dp[i][j][0] = min(dp[i][j-1][0]*grid[i][j], dp[i][j-1][1]*grid[i][j], dp[i-1][j][0]*grid[i][j], dp[i-1][j][1]*grid[i][j])
			dp[i][j][1] = max(dp[i][j-1][0]*grid[i][j], dp[i][j-1][1]*grid[i][j], dp[i-1][j][0]*grid[i][j], dp[i-1][j][1]*grid[i][j])
		}
	}
	if dp[r-1][c-1][1] >= 0 {
		return dp[r-1][c-1][1] % mod
	}
	return -1

}
