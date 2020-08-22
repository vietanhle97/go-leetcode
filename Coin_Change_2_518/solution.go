package Coin_Change_2_518

func change(amount int, coins []int) int {
	table := make([]int, amount+1)
	table[0] = 1

	for _, coin := range coins {
		for v := coin; v < amount+1; v++ {
			table[v] += table[v-coin]
		}
	}
	return table[amount]
}
