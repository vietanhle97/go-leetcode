package Best_Time_To_Buy_And_Sell_Stock_With_Transaction_Fee_714

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxProfit(P []int, fee int) int {
	m := len(P)
	if m == 1 {
		return 0
	}
	own, sell := -P[0], 0
	for i := 1; i < m; i++ {
		own = max(own, sell-P[i])
		sell = max(sell, own+P[i]-fee)

	}
	// fmt.Println(dp)
	return max(own, sell)
}
