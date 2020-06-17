package Sum_Of_Distances_In_Tree_834

func DFS(graph map[int][]int, node, parent int, count, ans *[]int) {
	for _, child := range graph[node] {
		if child != parent {
			DFS(graph, child, node, count, ans)
			(*count)[node] += (*count)[child]
			(*ans)[node] += (*ans)[child] + (*count)[child]
		}
	}
}

func DFS2(graph map[int][]int, node, parent, N int, count, ans *[]int) {
	for _, child := range graph[node] {
		if child != parent {
			(*ans)[child] = (*ans)[node] - (*count)[child] + N - (*count)[child]
			DFS2(graph, child, node, N, count, ans)
		}
	}
}

func sumOfDistancesInTree(N int, edges [][]int) []int {
	graph := map[int][]int{}
	for _, e := range edges {
		if _, ok := graph[e[0]]; !ok {
			graph[e[0]] = make([]int, 0)
		}
		if _, ok := graph[e[1]]; !ok {
			graph[e[1]] = make([]int, 0)
		}
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	count := make([]int, N)
	for i := range count {
		count[i] = 1
	}
	ans := make([]int, N)
	DFS(graph, 0, -1, &count, &ans)
	DFS2(graph, 0, -1, N, &count, &ans)
	return ans
}
