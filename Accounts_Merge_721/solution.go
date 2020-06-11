package Accounts_Merge_721

import (
	"sort"
)

func find(parent []int, i int) int {
	if parent[i] == i {
		return i
	}
	return find(parent, parent[i])
}

func union(parent []int, x, y int) {
	xSet := find(parent, x)
	ySet := find(parent, y)
	parent[xSet] = parent[ySet]
}

func accountsMerge(accounts [][]string) [][]string {
	parent := make([]int, 10001)
	for i := range parent {
		parent[i] = i
	}
	emToName := map[string]string{}
	emToId := map[string]int{}
	i := 0

	for _, acc := range accounts {
		name := acc[0]
		for _, em := range acc[1:] {
			emToName[em] = name
			if _, ok := emToId[em]; !ok {
				emToId[em] = i
				i += 1
			}
			union(parent, emToId[acc[1]], emToId[em])
		}
	}

	ans := map[int][]string{}
	for k, _ := range emToName {
		tmp := find(parent, emToId[k])
		if _, ok := ans[tmp]; !ok {
			ans[tmp] = []string{k}
		} else {
			ans[tmp] = append(ans[tmp], k)
		}
	}
	res := make([][]string, 0)

	for _, v := range ans {
		addr := []string{emToName[v[0]]}
		sort.Strings(v)
		addr = append(addr, v...)
		res = append(res, addr)
	}
	return res
}
