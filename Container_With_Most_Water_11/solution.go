package Container_With_Most_Water_11

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

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	res := 0
	for i < j {
		res = max(res, min(height[i], height[j])*(j-i))
		if height[j] < height[i] {
			j -= 1
		} else {
			i += 1
		}
	}
	return res
}
