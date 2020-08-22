package Sort_The_Matrix_Diagonally_1329

import "sort"

type p struct {
	i int
	j int
}

func diagonalSort(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	P := map[p][]int{}

	for i := 0; i < n; i++ {
		start, j, k := p{0, i}, 0, i
		for j < m && k < n {
			P[start] = append(P[start], mat[j][k])
			j++
			k++
		}
	}
	for i := 1; i < m; i++ {
		start, j, k := p{i, 0}, i, 0
		for j < m && k < n {
			P[start] = append(P[start], mat[j][k])
			j++
			k++
		}
	}
	for _, v := range P {
		sort.Ints(v)
	}
	for k, v := range P {
		x, y := k.i, k.j
		for _, e := range v {
			mat[x][y] = e
			x++
			y++
		}
	}
	return mat
}
