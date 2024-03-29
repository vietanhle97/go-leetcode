package Number_of_Dice_Rolls_With_Target_Sum_1155

func pow() int {
	res := 10
	for i := 1; i < 9; i++ {
		res *= 10
	}
	return res + 7
}

func numRollsToTarget(d int, f int, target int) int {
	if target < d {
		return 0
	}
	table := make([][]int, d+1)
	for i, _ := range table {
		table[i] = make([]int, target+1)
	}

	for i := 1; i < d+1; i++ {
		for j := i; j < target+1; j++ {
			if i == 1 {
				if j >= f+1 {
					break
				}
				if j <= target {
					table[i][j] = 1
				} else {
					break
				}

			} else if i == j {
				table[i][j] = 1
			} else {
				maxTarget := target - d + i
				minStart := 1
				for minStart <= j && j <= maxTarget {
					table[i][j] += table[i-1][minStart] * table[1][j-minStart]
					table[i][j] %= pow()
					minStart += 1
				}
			}
		}
	}
	return table[d][target]
}
