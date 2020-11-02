package Rank_Transform_Of_A_Matrix_1632

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func find(p []int, x int) int {
	if p[x] != x {
		p[x] = find(p, p[x])
	}
	return p[x]
}

func nrange(a int) []int {
	res := make([]int, 0)
	for i := 0; i < a; i++ {
		res = append(res, i)
	}
	return res
}

func matrixRankTransform(M [][]int) [][]int {
	m := len(M)
	n := len(M[0])
	rank := make([]int, m+n)
	d := map[int][][]int{}
	S := make([]int, 0)
	for i := range M {
		for j := range M[i] {
			if _, ok := d[M[i][j]]; !ok {
				S = append(S, M[i][j])
			}
			d[M[i][j]] = append(d[M[i][j]], []int{i, j})
		}
	}
	sort.Ints(S)
	for _, e := range S {
		p := nrange(m + n)
		rank2 := make([]int, m+n)
		copy(rank2, rank)
		for _, v := range d[e] {
			a, b := find(p, v[0]), find(p, v[1]+m)
			p[a] = b
			rank2[b] = max(rank2[a], rank2[b])
		}
		for _, v := range d[e] {
			tmp := rank2[find(p, v[0])] + 1
			rank[v[0]], rank[v[1]+m], M[v[0]][v[1]] = tmp, tmp, tmp
		}
	}
	return M
}
