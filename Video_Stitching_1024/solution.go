package Video_Stitching_1024

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func videoStitching(clips [][]int, T int) int {
	m := map[int][]int{}
	for _, e := range clips {
		for j := e[0]; j <= e[1] && j <= T; j++ {
			m[e[0]] = append(m[e[0]], j)
		}

	}
	if _, ok := m[0]; !ok {
		return -1
	}
	d := map[int]int{0: 0}
	vis := map[int]bool{}
	stack := []int{0}
	for len(stack) > 0 {
		cur := stack[0]
		stack = stack[1:]
		vis[cur] = true
		for _, e := range m[cur] {
			if _, ok := d[e]; !ok {
				d[e] = d[cur] + 1
			}
			d[e] = Min(d[e], d[cur]+1)
			if _, ok := m[e]; ok {
				if _, ok := vis[e]; !ok {
					stack = append(stack, e)
				}
			}
		}
	}
	if _, ok := d[T]; !ok {
		return -1
	}
	return d[T]
}
