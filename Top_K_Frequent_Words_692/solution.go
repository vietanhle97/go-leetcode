package Top_K_Frequent_Words_692

import "container/heap"

type Node struct {
	w string
	f int
}

type minHeap []*Node

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i, j int) bool {
	if h[i].f == h[j].f {
		return h[i].w < h[j].w
	}
	return h[i].f > h[j].f
}
func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	tmp := x.(Node)
	*h = append(*h, &tmp)
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	m := map[string]int{}
	for _, e := range words {
		m[e]++
	}
	h := &minHeap{}
	for k, v := range m {
		heap.Push(h, Node{k, v})
	}
	heap.Init(h)
	res := make([]string, 0)
	for len(res) < k {
		cur := heap.Pop(h).(*Node)
		res = append(res, cur.w)
	}
	return res
}
