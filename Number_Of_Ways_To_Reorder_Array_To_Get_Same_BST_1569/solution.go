package Number_Of_Ways_To_Reorder_Array_To_Get_Same_BST_1569

const mod = int(1e9) + 7

var dp [][]int

func dfs(A []int) int {
	n := len(A)
	if n <= 2 {
		return 1
	}
	L, R := make([]int, 0), make([]int, 0)
	for i := 1; i < n; i++ {
		if A[i] < A[0] {
			L = append(L, A[i])
		} else {
			R = append(R, A[i])
		}
	}
	lenLeft, resLeft, resRight := len(L), dfs(L), dfs(R)
	return (((dp[n-1][lenLeft] * resLeft) % mod) * resRight) % mod
}

func numOfWays(nums []int) int {
	m := len(nums)
	if m <= 2 {
		return 0
	}
	dp = make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			dp[i][j] = 1
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j < i; j++ {
			dp[i][j] = (dp[i-1][j-1] + dp[i-1][j]) % mod
		}
	}
	return dfs(nums)%mod - 1
}
