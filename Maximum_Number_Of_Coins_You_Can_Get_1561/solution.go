package Maximum_Number_Of_Coins_You_Can_Get_1561

import "sort"

func maxCoins(piles []int) int {
	sort.Ints(piles)
	m := len(piles)
	i, j, k := 0, m-2, m-1
	res := 0
	for i < j {
		res += piles[j]
		i++
		j -= 2
		k -= 2
	}
	return res
}
