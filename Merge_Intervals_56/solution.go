package Merge_Intervals_56

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func partition(low, high int, m [][]int) int {
	i := low - 1
	pivot := m[high]
	for j := low; j < high; j++ {
		if m[j][0] < pivot[0] {
			i += 1
			m[i], m[j] = m[j], m[i]
		}
	}
	m[i+1], m[high] = m[high], m[i+1]
	return i + 1
}

func quickSort(low, high int, m [][]int) {
	if low < high {
		pivot := partition(low, high, m)
		quickSort(low, pivot-1, m)
		quickSort(pivot+1, high, m)
	}
}

func merge(intervals [][]int) [][]int {
	quickSort(0, len(intervals)-1, intervals)
	output := make([][]int, 0)

	for _, e := range intervals {
		if len(output) == 0 || e[0] > output[len(output)-1][1] {
			output = append(output, e)
		} else {
			output[len(output)-1][1] = max(e[1], output[len(output)-1][1])
		}
	}
	return output
}
