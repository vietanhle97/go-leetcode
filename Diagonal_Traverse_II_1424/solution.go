package Diagonal_Traverse_II_1424

func findDiagonalOrder(nums [][]int) []int {
	m := len(nums)
	vis := make([][]bool, m)
	res := make([]int, 0)
	for i := range nums {
		vis[i] = make([]bool, len(nums[i]))
	}
	stack := [][]int{{0, 0}}

	for len(stack) > 0 {
		c := stack[0]
		stack = stack[1:]
		if vis[c[0]][c[1]] {
			continue
		}
		vis[c[0]][c[1]] = true
		res = append(res, nums[c[0]][c[1]])

		if c[0]+1 < m && len(nums[c[0]+1]) > c[1] {
			stack = append(stack, []int{c[0] + 1, c[1]})
		}

		if c[1]+1 < len(nums[c[0]]) {
			stack = append(stack, []int{c[0], c[1] + 1})
		}
	}
	return res
}
