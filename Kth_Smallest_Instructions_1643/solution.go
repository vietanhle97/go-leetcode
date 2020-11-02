package Kth_Smallest_Instructions_1643

func nCr() [][]int {
	res := make([][]int, 31)
	res[0] = []int{1}
	res[1] = []int{1, 1}
	for i := 2; i <= 30; i++ {
		res[i] = []int{1}
		for j := 1; j < i; j++ {
			res[i] = append(res[i], res[i-1][j-1]+res[i-1][j])
		}
		res[i] = append(res[i], 1)
	}
	return res
}
func fill(s string, tot, v int) string {
	for i := 0; i < v; i++ {
		s += "V"
	}
	for i := 0; i < tot-v; i++ {
		s += "H"
	}
	return s
}

func kthSmallestPath(des []int, k int) string {
	comb := nCr()
	tot, l := des[0]+des[1], des[0]+des[1]
	v := des[0]
	s := ""
	if k == comb[tot][v] {
		return fill(s, tot, v)
	}
	for len(s) < l {
		if k == comb[tot][v] {
			return fill(s, tot, v)
		} else if k > comb[tot-1][v] {
			s += "V"
			k -= comb[tot-1][v]
			v--
			tot--
		} else {
			s += "H"
			tot--
		}
	}
	return s
}
