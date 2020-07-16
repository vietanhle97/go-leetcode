package Minimum_Domino_Rotations_For_Equal_Row_1007

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minDominoRotations(A []int, B []int) int {
	a, b, same := [7]int{}, [7]int{}, [7]int{}
	m := len(A)
	for i := range A {
		if A[i] == B[i] {
			same[A[i]]++
		}
		a[A[i]]++
		b[B[i]]++
	}

	for i := range a {
		if a[i]+b[i]-same[i] == m {
			return m - max(a[i], b[i])
		}
	}
	return -1
}
