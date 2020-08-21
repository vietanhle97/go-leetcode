package Best_Time_To_Buy_And_Sell_Stock_IV_188

func max(a []int) int {
	res := a[0]
	for _, e := range a {
		if e > res {
			res = e
		}
	}
	return res
}
func maxProfit(k int, P []int) int {
	if k == 0 {
		return 0
	}
	A, n := make([]int, 0), len(P)
	if n <= 1 {
		return 0
	}
	for i := 0; i < n-1; i++ {
		A = append(A, P[i+1]-P[i])
	}
	if k > n/2 {
		res := 0
		for i := range A {
			if A[i] > 0 {
				res += A[i]
			}
		}
		return res
	}
	dp, mp := make([][]int, n-1), make([][]int, n-1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		mp[i] = make([]int, k+1)
	}
	dp[0][1], mp[0][1] = A[0], A[0]
	for i := 1; i < n-1; i++ {
		for j := 1; j < k+1; j++ {
			dp[i][j] = max([]int{mp[i-1][j-1], dp[i-1][j]}) + A[i]
			mp[i][j] = max([]int{mp[i-1][j], dp[i][j]})
		}
	}
	return max(mp[len(mp)-1])
}
