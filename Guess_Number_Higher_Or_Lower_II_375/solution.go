package Guess_Number_Higher_Or_Lower_II_375

const maxInt = 1<<63 - 1

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findCost(l, r int, table *[][]int) int {
	if (*table)[l][r] != -1 {
		return (*table)[l][r]
	} else if l == r {
		(*table)[l][r] = 0
		return 0
	} else if r == l+1 {
		(*table)[l][r] = l
		return l
	}
	cost := maxInt
	for mid := l + 1; mid < r; mid++ {
		ls := mid + findCost(l, mid-1, table)
		rs := mid + findCost(mid+1, r, table)
		if max(ls, rs) < cost {
			cost = max(ls, rs)
		}
	}
	(*table)[l][r] = cost
	return cost
}

func getMoneyAmount(n int) int {
	table := make([][]int, n+1)
	for i := range table {
		table[i] = make([]int, n+1)
		for j := range table[i] {
			table[i][j] = -1
		}
	}
	findCost(1, n, &table)
	return table[1][n]
}
