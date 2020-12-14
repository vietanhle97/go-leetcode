package Minimum_Jumps_To_Reach_Home_1654

type node struct {
	i    int
	prev int
	cnt  int
}

func check(a []int) (map[int]int, int) {
	m, res := map[int]int{}, 0
	for _, e := range a {
		m[e] = 1
		if e > res {
			res = e
		}
	}
	return m, res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minimumJumps(f []int, a int, b int, x int) int {
	m, maxf := check(f)
	if _, ok := m[a]; ok {
		return -1
	}
	if x == 0 {
		return 0
	}
	maxVal := max(x, maxf) + a + b
	q := make([]node, 0)
	q = append(q, node{a, 1, 1})
	vis := map[node]int{}
	res := -1
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if vis[node{cur.i, cur.prev, 1}] == 1 {
			continue
		}
		vis[node{cur.i, cur.prev, 1}] = 1
		if cur.i == x {
			res = cur.cnt
			break
		}
		if _, ok := m[cur.i+a]; !ok && cur.i+a-b <= maxVal {
			q = append(q, node{cur.i + a, 1, cur.cnt + 1})
		}
		if _, ok := m[cur.i-b]; !ok && cur.prev == 1 && cur.i-b >= 0 {
			q = append(q, node{cur.i - b, -1, cur.cnt + 1})
		}
	}
	return res
}
