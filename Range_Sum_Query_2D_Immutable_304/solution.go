package Range_Sum_Query_2D_Immutable_304

type NumMatrix struct {
	mat [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	mat := make([][]int, m)
	for i := range matrix {
		for j := range matrix[i] {
			if j == 0 {
				mat[i] = append(mat[i], 0)
			}
			mat[i] = append(mat[i], matrix[i][j]+mat[i][len(mat[i])-1])
		}
	}
	return NumMatrix{mat}
}

func (m *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	res := 0
	for i := row1; i <= row2; i++ {
		res += m.mat[i][col2+1] - m.mat[i][col1]
	}
	return res
}
