package Search_A_2D_Matrix_II_240

func binarySearch(nums []int, l, r, target int) bool {
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	cnt := -1
	for i := range matrix {
		if target <= matrix[i][n-1] {
			if target == matrix[i][n-1] {
				return true
			}
			cnt = i
			break
		}
	}
	if cnt == -1 {
		return false
	}
	for cnt < m {
		if binarySearch(matrix[cnt], 0, n-1, target) {
			return true
		}
		cnt++
	}
	return false
}
