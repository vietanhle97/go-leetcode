package Course_Schedule_II_210

func isExist(nums []int, i int) (bool, int) {
	for _, e := range nums {
		if e == i {
			return true, -1
		}
	}
	return false, i
}

func DFS(graph map[int][]int, v int, visited, rec []bool, stack *[]int) bool {
	visited[v], rec[v] = true, true
	for _, neigh := range graph[v] {
		if visited[neigh] == false {
			if DFS(graph, neigh, visited, rec, stack) {
				return true
			}
		}
		if rec[neigh] {
			return true
		}
	}
	rec[v] = false
	*stack = append(*stack, v)
	return false
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	if len(prerequisites) == 0 {
		res := make([]int, numCourses)
		for i := 0; i < numCourses; i++ {
			res[i] = i
		}
		return res
	}
	visited, rec := make([]bool, numCourses), make([]bool, numCourses)
	graph := map[int][]int{}
	stack := make([]int, 0)
	for _, e := range prerequisites {
		if _, ok := graph[e[0]]; !ok {
			graph[e[0]] = []int{e[1]}
		} else {
			graph[e[0]] = append(graph[e[0]], e[1])
		}
	}

	for v, _ := range graph {
		if visited[v] == false {
			if DFS(graph, v, visited, rec, &stack) {
				return []int{}
			}
		}
	}
	if len(stack) < numCourses {
		for i := 0; i < numCourses; i++ {
			b, e := isExist(stack, i)
			if !b {
				stack = append(stack, e)
			}
		}
	}
	return stack
}
