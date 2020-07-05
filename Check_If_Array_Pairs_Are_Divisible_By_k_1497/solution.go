package Check_If_Array_Pairs_Are_Divisible_By_k_1497

func canArrange(arr []int, k int) bool {
	m := make([]int, k)
	sum := 0
	for _, e := range arr {
		tmp := e % k
		if tmp < 0 {
			tmp += k
		}
		if tmp == 0 {
			if m[tmp] == 0 {
				m[tmp]++
				sum++
			} else {
				m[tmp]--
				sum--
			}
		} else {
			if m[k-tmp] > 0 {
				m[k-tmp]--
				sum--
			} else {
				m[tmp]++
				sum++
			}
		}
	}
	return sum == 0
}
