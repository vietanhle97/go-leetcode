package Maximum_Sum_Of_Two_Non_Overlapping_Subarrays_1031

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxSumTwoNoOverlap(A []int, L int, M int) int {
	sum, n := 0, len(A)
	if L+M == n {
		for _, e := range A {
			sum += e
		}
		return sum
	}

	L1 := make([]int, 0)
	M1 := make([]int, 0)
	l1, m1 := 0, 0
	for i, j := 0, 0; i < L || j < M; i, j = i+1, j+1 {
		if i < L {
			l1 += A[i]
		}
		if j < M {
			m1 += A[j]
		}
	}
	L1, M1 = append(L1, l1), append(M1, m1)
	for i, j := L, M; i < n || j < n; i, j = i+1, j+1 {
		if i < n {
			L1 = append(L1, L1[len(L1)-1]-A[i-L]+A[i])
		}
		if j < n {
			M1 = append(M1, M1[len(M1)-1]-A[j-M]+A[j])
		}
	}
	res := 0
	for i, e1 := range L1 {
		for j, e2 := range M1 {
			if i+L <= j || j+M <= i {
				res = max(res, e1+e2)
			}
		}
	}

	return res
}
