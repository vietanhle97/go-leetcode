package Furthest_Building_You_Can_Reach_1642

import "container/heap"

type minHeap []int

func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h minHeap) Len() int           { return len(h) }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func furthestBuilding(H []int, bricks int, ladders int) int {
	m := len(H)
	h := &minHeap{}
	heap.Init(h)
	cur := 0
	for i := 1; i < m; i++ {
		if H[i] > H[i-1] {
			heap.Push(h, H[i]-H[i-1])
			if len(*h) > ladders {
				min := heap.Pop(h).(int)
				if bricks >= min {
					bricks -= min
					cur = i
				} else {
					break
				}
			} else {
				cur = i
			}
		} else {
			cur = i
		}
	}
	return cur
}
