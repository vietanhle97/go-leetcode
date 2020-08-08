package Coin_Change_2_518

func change(amount int, coins []int) int {
	table := make([]int, amount+1)
	table[0] = 1

	for c := 0; c < len(coins); c++ {
		for v := coins[c]; v < amount+1; v++ {
			table[v] += table[v-coins[c]]
		}
	}
	return table[amount]
}
