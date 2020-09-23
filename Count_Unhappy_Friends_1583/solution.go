package Count_Unhappy_Friends_1583

func check(pairs [][]int, P []map[int]int, x, y, i int) int {
	for j, newPair := range pairs {
		if i == j {
			continue
		}
		u, v := newPair[0], newPair[1]
		if P[x][u] < P[x][y] && P[u][x] < P[u][v] {
			return 1
		}
		u, v = newPair[1], newPair[0]
		if P[x][u] < P[x][y] && P[u][x] < P[u][v] {
			return 1
		}
	}
	return 0
}

func unhappyFriends(n int, p [][]int, pairs [][]int) int {
	P := make([]map[int]int, n)
	for i := range P {
		P[i] = map[int]int{}
		for j, e := range p[i] {
			P[i][e] = j
		}
	}
	cnt := 0
	for i, pair := range pairs {
		x, y := pair[0], pair[1]
		if p[x][0] != y {
			cnt += check(pairs, P, x, y, i)
		}
		if p[y][0] != x {
			cnt += check(pairs, P, y, x, i)
		}
	}
	return cnt
}
