package Rotate_Image_48

func rotate(A [][]int) {
	n := len(A)
	for i := 0; i < n; i++ {
		for j := i; j < n-1-i; j++ {
			A[i][j], A[j][n-1-i], A[n-1-i][n-1-j], A[n-1-j][i] = A[n-1-j][i], A[i][j], A[j][n-1-i], A[n-1-i][n-1-j]
		}
	}
}
