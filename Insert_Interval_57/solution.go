package Insert_Interval_57

import (
	"math"
)

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func merge(output *[][]int, interval []int) {
	if len(*output) == 0 || interval[0] > (*output)[len(*output)-1][1] {
		*output = append(*output, interval)
	} else {
		(*output)[len(*output)-1][1] = max(interval[1], (*output)[len(*output)-1][1])
	}
}

func insert(intervals [][]int, newInterval []int) [][]int {
	ind := -1

	for i, e := range intervals {
		if e[0] == newInterval[0] {
			if e[1] < newInterval[1] {
				e[1] = newInterval[1]
				break
			} else {
				return intervals
			}
		}
		if e[0] > newInterval[0] {
			ind = i
			break
		}
	}
	if ind == -1 {
		intervals = append(intervals, newInterval)
	}
	output := make([][]int, 0)
	i := 0
	check := false
	for i < len(intervals) {
		if i == ind {
			if !check {
				merge(&output, newInterval)
				check = true
			} else {
				merge(&output, intervals[i])
				i++
			}
		} else {
			merge(&output, intervals[i])
			i++
		}
	}
	return output
}
