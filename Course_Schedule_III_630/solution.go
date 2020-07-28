package Course_Schedule_III_630

import (
	"container/heap"
	"sort"
)

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]
	return res
}

func scheduleCourse(c [][]int) int {
	sort.Slice(c, func(i, j int) bool {
		if c[i][1] == c[j][1] {
			return c[i][0] < c[j][0]
		}
		return c[i][1] < c[j][1]
	})
	time, take := 0, &maxHeap{}
	heap.Init(take)
	for i := range c {
		if c[i][0]+time <= c[i][1] {
			time += c[i][0]
			heap.Push(take, c[i][0])
		} else {
			if len(*take) > 0 && (*take)[0] > c[i][0] {
				time += c[i][0] - heap.Pop(take).(int)
				heap.Push(take, c[i][0])
			}
		}
	}
	return len(*take)
}
