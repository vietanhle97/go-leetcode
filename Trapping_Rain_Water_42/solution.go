package Trapping_Rain_Water_42

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trap(height []int) int {
	m := len(height)
	left := make([]int, m)
	right := make([]int, m)
	cnt := 0
	for i := range height {
		if i == 0 {
			left[i] = height[i]
			right[m-1-i] = height[m-1-i]
		} else {
			left[i] = max(height[i], left[i-1])
			right[m-1-i] = max(height[m-1-i], right[m-i])
		}
	}

	for i := range height {
		cnt += min(left[i]-height[i], right[i]-height[i])
	}
	return cnt
}
