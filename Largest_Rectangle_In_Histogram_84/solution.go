package Largest_Rectangle_In_Histogram_84

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(h []int) int {
	m := len(h)
	if m == 0 {
		return 0
	}
	res := 0
	for i := 0; i < m; i++ {
		tmp := h[i]
		for j := i; j < m; j++ {
			if h[j] < tmp {
				tmp = h[j]
			}
			res = max(res, tmp*(j-i+1))
		}
	}
	return res
}
