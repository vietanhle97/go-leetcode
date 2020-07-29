package Merge_Intervals_56

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
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
