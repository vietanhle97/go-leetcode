package Minimum_Swaps_To_Arrange_A_Binary_Grid_1536

func minSwaps(grid [][]int) int {
	n := len(grid)
	v := make([]int, n)
	for i := 0; i < n; i++ {
		tmp, last := grid[i], -1
		for j, e := range tmp {
			if e == 1 {
				last = j
			}
		}
		v[i] = last
	}
	ans := 0
	for i := 0; i < n; i++ {
		valid := false
		for j, e := range v {
			if e <= i {
				ans += j
				v = append(v[:j], v[j+1:]...)
				valid = true
				break
			}
		}
		if !valid {
			return -1
		}
	}
	return ans
}
