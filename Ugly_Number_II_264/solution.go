package Ugly_Number_II_264

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func nthUglyNumber(n int) int {
	mul := []int{2, 3, 5}
	h := &IntHeap{2, 3, 5}
	check := map[int]bool{2: true, 3: true, 5: true}
	heap.Init(h)
	for n > 0 {
		cur := heap.Pop(h)
		n -= 1
		if n == 1 {
			return cur.(int)
		}
		for _, e := range mul {
			if _, ok := check[cur.(int)*e]; !ok {
				heap.Push(h, cur.(int)*e)
				check[cur.(int)*e] = true
			}
		}
	}
	return 1
}
