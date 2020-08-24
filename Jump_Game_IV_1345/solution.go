package Jump_Game_IV_1345

func minJumps(A []int) int {
	n, m := len(A), map[int][]int{}
	for i, e := range A {
		m[e] = append(m[e], i)
	}
	q, numVis, vis, dist := []int{0}, map[int]bool{}, make([]bool, n), 0
	vis[0] = true
	for len(q) > 0 {
		tmp := len(q)
		for _, e := range q {
			if e == n-1 {
				return dist
			}
			if !numVis[A[e]] {
				numVis[A[e]] = true
				for _, v := range m[A[e]] {
					if v == e {
						continue
					}
					q = append(q, v)
					vis[v] = true
				}
			}
			if e-1 >= 0 && !vis[e-1] {
				q = append(q, e-1)
				vis[e-1] = true
			}
			if e+1 < n && !vis[e+1] {
				q = append(q, e+1)
				vis[e+1] = true
			}
		}
		q = q[tmp:]
		dist++
	}
	return -1
}
