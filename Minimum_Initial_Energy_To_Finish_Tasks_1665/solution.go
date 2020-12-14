package Minimum_Initial_Energy_To_Finish_Tasks_1665

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minimumEffort(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1]-tasks[i][0] < tasks[j][1]-tasks[j][0]
	})
	res := 0
	for _, e := range tasks {
		res = max(res+e[0], e[1])
	}
	return res
}
