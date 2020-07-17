package Time_Needed_To_Inform_All_Employees_1376

import "math"

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func dfs(v int, g map[int][]int, informTime []int, dist int, res *int) {
	for _, e := range g[v] {
		d := dist + informTime[e]
		*res = max(*res, d)
		dfs(e, g, informTime, d, res)
	}
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	g := map[int][]int{}
	res := 0
	for i, e := range manager {
		if e != -1 {
			g[e] = append(g[e], i)
		}
	}
	dfs(headID, g, informTime, informTime[headID], &res)
	return res

}
