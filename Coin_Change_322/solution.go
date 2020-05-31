package Coin_Change_322

func isExist(coins []int, n int) bool {
	for _, e := range coins {
		if e == n {
			return true
		}
	}
	return false
}

func min_(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	res := []int{}
	for i := 0; i < amount+1; i++ {
		if isExist(coins, i) {
			res = append(res, 1)
		} else {
			res = append(res, -1)
		}
	}
	for i := 1; i < amount+1; i++ {
		for _, e := range coins {
			if e < i {
				if res[i-e] == -1 {
					continue
				} else {
					if res[i] == -1 {
						res[i] = res[i-e] + 1
					} else {
						res[i] = min_(res[i], res[i-e]+1)
					}

				}
			}
		}
	}
	return res[amount]
}
