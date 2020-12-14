package Max_Number_Of_K_Sum_Pairs_1679

func maxOperations(nums []int, k int) int {
	m := map[int]int{}
	cnt := 0
	for _, e := range nums {
		if _, ok := m[e]; !ok || m[e] == 0 {
			m[k-e]++
		} else {
			m[e]--
			cnt++
		}
	}
	return cnt
}
