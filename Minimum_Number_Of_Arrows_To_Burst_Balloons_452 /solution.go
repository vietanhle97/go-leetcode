package Minimum_Number_Of_Arrows_To_Burst_Balloons_452_

import "sort"

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

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	prev, cnt := []int{-1, -1}, 0
	for _, e := range points {
		if (prev[0] == -1 && prev[1] == -1) || e[0] > prev[1] {
			prev = e
			cnt++
		} else {
			prev[0] = max(e[0], prev[0])
			prev[1] = min(e[1], prev[1])
		}
	}
	return cnt
}
