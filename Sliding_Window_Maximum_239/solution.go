package Sliding_Window_Maximum_239

// using heap -> complexity is O(NlogN)
import "container/heap"

type maxHeap [][]int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i][0] > h[j][0] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *maxHeap) Pop() interface{} {
	n := len(*h) - 1
	res := (*h)[n]
	*h = (*h)[:n]
	return res
}

func maxSlidingWindow(nums []int, k int) []int {
	h, n, res, cnt := &maxHeap{}, len(nums), make([]int, 0), 0
	heap.Init(h)
	for i := 0; i < k; i++ {
		heap.Push(h, []int{nums[i], i})
	}
	for i := k; i <= n; i++ {
		cur := heap.Pop(h).([]int)
		for cur[1] < cnt {
			cur = heap.Pop(h).([]int)
		}
		res = append(res, cur[0])
		if cur[1] > cnt {
			heap.Push(h, cur)
		}
		if i < n {
			heap.Push(h, []int{nums[i], i})
		}
		cnt++
	}

	return res
}
