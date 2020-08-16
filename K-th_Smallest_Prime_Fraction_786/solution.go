package K_th_Smallest_Prime_Fraction_786

import "container/heap"

type minHeap [][]int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i][0]*h[j][1] < h[i][1]*h[j][0] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]
	return res
}

func kthSmallestPrimeFraction(A []int, K int) []int {
	h, n := &minHeap{}, len(A)-1
	heap.Init(h)
	heap.Push(h, []int{A[0], A[n], 0, n})
	cnt, vis := 0, make([][]int, n+1)
	for i := range A {
		vis[i] = make([]int, n+1)
	}
	i, j := 0, 0
	for cnt < K {
		cur := heap.Pop(h).([]int)
		i, j = cur[2], cur[3]
		cnt++
		if j > 0 && vis[i][j-1] == 0 {
			heap.Push(h, []int{A[i], A[j-1], i, j - 1})
			vis[i][j-1] = 1
		}
		if i < n && vis[i+1][j] == 0 {
			heap.Push(h, []int{A[i+1], A[j], i + 1, j})
			vis[i+1][j] = 1
		}
		if cnt == K {
			break
		}
	}
	return []int{A[i], A[j]}
}
