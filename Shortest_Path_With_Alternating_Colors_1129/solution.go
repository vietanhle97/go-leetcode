package Shortest_Path_With_Alternating_Colors_1129

import "strconv"

func shortestAlternatingPaths(n int, red_edges [][]int, blue_edges [][]int) []int {
	res := []int{0}
	for i := 0; i < n-1; i++ {
		res = append(res, -1)
	}
	if len(red_edges) == 0 && len(blue_edges) == 0 {
		return res
	}

	graph := map[int][][]int{}
	stack := make([][]int, 0)
	dist := map[int]int{}
	visited := map[string]bool{}

	for _, e := range red_edges {
		graph[e[0]] = append(graph[e[0]], append([]int{e[1], 0}))
	}
	for _, e := range blue_edges {
		graph[e[0]] = append(graph[e[0]], append([]int{e[1], 1}))
	}
	if _, ok := graph[0]; !ok {
		return res
	}
	for i := 0; i < n; i++ {
		dist[i] = -1
	}
	dist[0] = 0
	stack = append(stack, []int{0, 0, 0})
	stack = append(stack, []int{0, 1, 0})
	for len(stack) > 0 {
		v, e, d := stack[0][0], stack[0][1], stack[0][2]
		stack = stack[1:]
		if v == 0 {
			for _, i := range graph[v] {
				tmp := strconv.Itoa(i[0]) + "," + strconv.Itoa(i[1])
				if _, ok := visited[tmp]; !ok {
					dist[i[0]] = 1
					stack = append(stack, []int{i[0], i[1], 1})
					visited[tmp] = true
				}
			}
		} else {
			if _, ok := graph[v]; ok {
				for _, i := range graph[v] {
					tmp := strconv.Itoa(i[0]) + "," + strconv.Itoa(i[1])
					if _, ok := visited[tmp]; !ok {
						if i[1] == 1-e {
							stack = append(stack, []int{i[0], i[1], d + 1})
							visited[tmp] = true
							if dist[i[0]] == -1 {
								dist[i[0]] = d + 1
							}
						}
					}
				}
			}
		}
	}
	for val, _ := range dist {
		if val != 0 {
			res[val] = dist[val]
		}
	}
	return res

}
