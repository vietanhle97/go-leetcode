package Number_Of_Ways_Where_Square_Of_Number_Is_Equal_To_Product_Of_Two_Numbers_1577

func numTriplets(A []int, B []int) int {
	m1, m2 := map[int]int{}, map[int]int{}
	n1, n2 := len(A), len(B)
	ans := 0
	for i := 0; i < n1-1; i++ {
		for j := i + 1; j < n1; j++ {
			m1[A[i]*A[j]]++
		}
	}

	for i := 0; i < n2-1; i++ {
		for j := i + 1; j < n2; j++ {
			m2[B[i]*B[j]]++
		}
	}

	for i := range A {
		if _, ok := m2[A[i]*A[i]]; ok {
			ans += m2[A[i]*A[i]]
		}
	}
	for i := range B {
		if _, ok := m1[B[i]*B[i]]; ok {
			ans += m1[B[i]*B[i]]
		}
	}
	return ans
}
