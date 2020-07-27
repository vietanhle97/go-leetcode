package Find_The_City_With_The_Smallest_Number_Of_Neighbors_At_A_Threshold_Distance_1334

import "container/heap"

const MaxInt = 1<<31 - 1

type node struct {
	v  int
	ed int
}
type minHeap []node

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].ed < h[j].ed }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(node)) }
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dijkstra(u, n, threshold int, g [][][]int, vis []bool) int {
	cnt := 0
	stack := &minHeap{}
	heap.Init(stack)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = MaxInt
	}
	dist[u], vis[u] = 0, true
	heap.Push(stack, node{u, 0})
	for stack.Len() > 0 {
		cur := heap.Pop(stack).(node)
		if cur.v != u && !vis[cur.v] {
			cnt++
		}
		vis[cur.v] = true
		for _, e := range g[cur.v] {
			dist[e[0]] = min(dist[e[0]], dist[cur.v]+e[1])
			if !vis[e[0]] && dist[e[0]] <= threshold {
				heap.Push(stack, node{e[0], dist[cur.v] + e[1]})
			}
		}
	}
	return cnt
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	g := make([][][]int, n)

	for _, e := range edges {
		g[e[0]] = append(g[e[0]], []int{e[1], e[2]})
		g[e[1]] = append(g[e[1]], []int{e[0], e[2]})
	}
	ans := n
	v := 0
	for i := 0; i < n; i++ {
		vis := make([]bool, n)
		tmp := dijkstra(i, n, distanceThreshold, g, vis)
		if tmp <= ans {
			ans = tmp
			v = max(v, i)
		}
	}
	return v
}
