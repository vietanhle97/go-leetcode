package Partition_Labels_763

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func partitionLabels(S string) []int {
	m := make([][]int, 26)
	for i := range S {
		if len(m[S[i]-'a']) <= 1 {
			m[S[i]-'a'] = append(m[S[i]-'a'], i)
		} else {
			m[S[i]-'a'][1] = i
		}
	}
	res := make([][]int, 0)
	for i := range m {
		if len(m[i]) > 0 {
			res = append(res, m[i])
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
	st := make([][]int, 0)
	for i := range res {
		if len(st) == 0 || len(st[len(st)-1]) == 1 {
			st = append(st, res[i])
		} else {
			if res[i][0] < st[len(st)-1][1] {
				if len(res[i]) > 1 {
					st[len(st)-1][1] = max(res[i][1], st[len(st)-1][1])
				}
			} else {
				st = append(st, res[i])
			}
		}
	}
	ans := make([]int, 0)
	for _, e := range st {
		if len(e) > 1 {
			ans = append(ans, e[1]-e[0]+1)
		} else {
			ans = append(ans, 1)
		}
	}
	return ans
}
