package Path_With_Maximum_Probability_1514

import (
	"container/heap"
	"math"
)

type node struct {
	v    int
	prob float64
}

type maxHeap []*node

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].prob > h[j].prob }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
	tmp := x.(*node)
	*h = append(*h, tmp)
}

func (h *maxHeap) Pop() interface{} {
	n := len(*h)
	tmp := (*h)[n-1]
	*h = (*h)[:n-1]
	return tmp
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	g, P := map[int][]node{}, map[int]float64{}
	P[start] = 1.0
	st := &node{start, 1.0}
	h := &maxHeap{st}
	heap.Init(h)
	for i := 0; i < n; i++ {
		g[i], P[i] = make([]node, 0), 0.0
	}
	for i, e := range edges {
		g[e[0]], g[e[1]] = append(g[e[0]], node{e[1], succProb[i]}), append(g[e[1]], node{e[0], succProb[i]})
	}
	for h.Len() > 0 {
		cur := heap.Pop(h).(*node)
		v, prob := cur.v, cur.prob
		for i, e := range g[v] {
			d := prob * e.prob
			if d > P[e.v] {
				P[e.v] = d
				tmp := g[v][i]
				tmp.prob = d
				heap.Push(h, &tmp)
			}
		}
	}
	return P[end]
}
