package Find_Latest_Group_Of_Size_M_1562

func findLatestStep(A []int, m int) int {
	n := len(A)
	length, cnt, res := make([]int, n+2), make([]int, n+1), -1
	for i, a := range A {
		left, right := length[a-1], length[a+1]
		length[a-left], length[a], length[a+right] = left+right+1, left+right+1, left+right+1
		cnt[left]--
		cnt[right]--
		cnt[length[a]]++
		if cnt[m] > 0 {
			res = i + 1
		}
	}
	return res
}
