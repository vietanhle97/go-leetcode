package Arithmetic_Slices_413

func numberOfArithmeticSlices(A []int) int {
	m := len(A)
	res := 0
	if m <= 2 {
		return res
	}
	for i := 1; i < m; i++ {
		prev, cur := A[i-1], A[i]
		d := cur - prev
		for j := i + 1; j < m; j++ {
			if A[j]-cur == d {
				res += 1
				cur = A[j]
			} else {
				break
			}
		}
	}
	return res
}
