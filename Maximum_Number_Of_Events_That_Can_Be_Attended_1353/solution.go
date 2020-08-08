package Maximum_Number_Of_Events_That_Can_Be_Attended_1353

import (
	"container/heap"
	"sort"
)

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]
	return res
}

func maxEvents(A [][]int) int {
	sort.Slice(A, func(i, j int) bool {
		if A[i][0] == A[j][0] {
			return A[i][1] < A[j][1]
		}
		return A[i][0] < A[j][0]
	})
	res, d, h := 0, 0, &minHeap{}
	heap.Init(h)
	for len(A) > 0 || h.Len() > 0 {
		if h.Len() == 0 {
			d = A[0][0]
		}
		for len(A) > 0 && A[0][0] == d {
			heap.Push(h, A[0][1])
			A = A[1:]
		}
		heap.Pop(h)
		res++
		d++
		for h.Len() > 0 && (*h)[0] < d {
			heap.Pop(h)
		}
	}
	return res
}
