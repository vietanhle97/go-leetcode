package Kth_Smallest_Element_In_A_Sorted_Matrix_378

import "container/heap"

type Node struct {
	val int
	r   int
	c   int
}

type minHeap []*Node

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	h := &minHeap{}
	cnt := 0
	for i := 0; i < n; i++ {
		heap.Push(h, Node{matrix[0][i], 0, i})
	}
	heap.Init(h)
	for cnt < k {
		tmp := heap.Pop(h).(*Node)
		cnt++
		if cnt == k {
			return tmp.val
		}
		if tmp.r+1 < n {
			heap.Push(h, Node{matrix[tmp.r+1][tmp.c], tmp.r + 1, tmp.c})
		}
	}
	return -1
}
