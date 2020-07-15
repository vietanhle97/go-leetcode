package Minimum_Insertion_Steps_To_Make_A_String_Palindrome_1312

func min(a ...int) int {
	res := a[0]
	for _, e := range a {
		if e < res {
			res = e
		}
	}
	return res
}
func minInsertions(s string) int {
	m := len(s)
	if m == 1 {
		return 0
	}
	ssp := make([][]int, m)
	for i := 0; i < m; i++ {
		ssp[i] = make([]int, m)
		ssp[i][i] = 1
	}
	for cl := 2; cl < m+1; cl++ {
		for i := 0; i < m-cl+1; i++ {
			j := i + cl - 1
			if s[i] == s[j] && cl == 2 {
				ssp[i][j] = 2
			} else if s[i] == s[j] {
				ssp[i][j] = 2 + min(ssp[i+1][j], ssp[i][j-1], ssp[i+1][j-1])
			} else {
				ssp[i][j] = 2 + min(ssp[i+1][j], ssp[i][j-1])

			}
		}
	}
	return ssp[0][m-1] - m

}
