package Merge_K_Sorted_Lists_23

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type minHeap []*ListNode

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]
	return res
}
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	res := &ListNode{}
	head := res
	h := &minHeap{}
	for _, e := range lists {
		if e != nil {
			heap.Push(h, e)
		}
	}
	heap.Init(h)
	for len(*h) > 0 {
		cur := heap.Pop(h).(*ListNode)
		head.Next = cur
		if cur.Next != nil {
			heap.Push(h, cur.Next)
		}
		head = head.Next
	}
	return res.Next
}
