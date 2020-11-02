package Maximal_Network_Rank_1615

type Node struct {
	u int
	v int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximalNetworkRank(n int, roads [][]int) int {
	m := make([]int, n)
	vis := map[Node]bool{}
	for _, e := range roads {
		m[e[0]]++
		m[e[1]]++
		vis[Node{e[0], e[1]}] = true
	}
	res := -(1<<31 - 1) - 1
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			tmp := max(res, m[i]+m[j])
			if vis[Node{i, j}] || vis[Node{j, i}] {
				tmp--
			}
			res = max(res, tmp)
		}
	}
	return res
}
