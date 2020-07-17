package Maximum_Frequency_Stack_895

import "container/heap"

type Node struct {
	cnt int
	ind int
	x   int
}

type maxHeap []*Node

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i, j int) bool {
	if h[i].cnt == h[j].cnt {
		return h[i].ind > h[j].ind
	}
	return h[i].cnt > h[j].cnt
}
func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}
func (h *maxHeap) Pop() interface{} {
	n := len(*h) - 1
	res := (*h)[n]
	*h = (*h)[:n]
	return res
}

type FreqStack struct {
	cnt   map[int]int
	stack *maxHeap
	ind   int
}

func Constructor() FreqStack {
	res := FreqStack{}
	res.ind, res.stack, res.cnt = -1, &maxHeap{}, map[int]int{}
	heap.Init(res.stack)
	return res
}

func (st *FreqStack) Push(x int) {
	st.ind++
	st.cnt[x]++
	e := Node{st.cnt[x], st.ind, x}
	heap.Push(st.stack, &e)

}

func (st *FreqStack) Pop() int {
	e := heap.Pop(st.stack).(*Node)
	st.cnt[e.x]--
	return e.x
}
