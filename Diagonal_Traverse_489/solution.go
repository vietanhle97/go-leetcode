package Diagonal_Traverse_489

func findDiagonalOrder(matrix [][]int) []int {
	res := make([]int, 0)
	m := len(matrix)
	if m == 0 {
		return res
	}
	n := len(matrix[0])
	if n == 0 {
		return res
	}
	up, cnt, cur := 1, m*n, []int{0, 0}
	for cnt > 0 {
		res = append(res, matrix[cur[0]][cur[1]])
		cnt--
		if up == 1 {
			if cur[0] > 0 && cur[1] < n-1 {
				cur[0]--
				cur[1]++
			} else if cur[1] == n-1 {
				cur[0]++
				up = 0
			} else if cur[0] == 0 {
				cur[1]++
				up = 0
			}
		} else {
			if cur[0] < m-1 && cur[1] > 0 {
				cur[0]++
				cur[1]--
			} else if cur[0] == m-1 {
				cur[1]++
				up = 1
			} else if cur[1] == 0 {
				cur[0]++
				up = 1
			}
		}
	}
	return res
}
