package Course_Schdule_207

func dfsCycle(m map[int][]int, visited []bool, rec []bool, v int) bool {
	visited[v] = true
	rec[v] = true
	for _, e := range m[v] {
		if visited[e] == false {
			if dfsCycle(m, visited, rec, e) {
				return true
			}
		}
		if rec[e] {
			return true
		}
	}
	rec[v] = false
	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	m := map[int][]int{}
	for _, e := range prerequisites {
		_, ok := m[e[0]]
		if ok {
			m[e[0]] = append(m[e[0]], e[1])
		} else {
			m[e[0]] = []int{e[1]}
		}
	}
	visited := make([]bool, numCourses)
	rec := make([]bool, numCourses)
	for k, _ := range m {
		if visited[k] == false {
			if dfsCycle(m, visited, rec, k) {
				return false
			}
		}
	}
	return true
}
