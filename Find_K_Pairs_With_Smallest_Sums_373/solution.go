package Find_K_Pairs_With_Smallest_Sums_373

import "container/heap"

type p struct {
	sum int
	x   int
	y   int
}

type minHeap []p

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i, j int) bool {
	if h[i].sum == h[j].sum {
		return h[i].x < h[j].x
	}
	return h[i].sum < h[j].sum
}
func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(p))
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]
	return res
}

func push(i, j int, A, B []int, h *minHeap) {
	if i < len(A) && j < len(B) {
		heap.Push(h, p{A[i] + B[j], i, j})
	}
}

func kSmallestPairs(A, B []int, k int) [][]int {
	res := make([][]int, 0)
	h := &minHeap{}
	heap.Init(h)
	push(0, 0, A, B, h)
	for len(*h) > 0 && len(res) < k {
		cur := heap.Pop(h).(p)
		x, y := cur.x, cur.y
		res = append(res, []int{A[x], B[y]})
		push(x, y+1, A, B, h)
		if y == 0 {
			push(x+1, y, A, B, h)
		}
	}
	return res
}
