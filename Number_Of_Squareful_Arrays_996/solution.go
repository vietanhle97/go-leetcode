package Number_Of_Squareful_Arrays_996

import (
	"math"
	"sort"
)

func DFS(nums, path []int, res *int) {
	if len(nums) == 0 {
		*res++
		return
	}
	for i, e := range nums {
		if len(path) == 0 {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			newNum := make([]int, i)
			copy(newNum, nums[:i])
			DFS(append(newNum, nums[i+1:]...), append(path, e), res)
		} else {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			square := math.Sqrt(float64(path[len(path)-1] + e))
			if square == float64(int(square)) {
				newNum := make([]int, i)
				copy(newNum, nums[:i])
				DFS(append(newNum, nums[i+1:]...), append(path, e), res)
			}
		}
	}
}

func numSquarefulPerms(A []int) int {
	res := 0
	sort.Ints(A)
	DFS(A, []int{}, &res)
	return res
}
