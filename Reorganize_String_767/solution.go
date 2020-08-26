package Reorganize_String_767

import "container/heap"

type char struct {
	ch rune
	c  int
}

type charHeap []char

func (h charHeap) Len() int           { return len(h) }
func (h charHeap) Less(i, j int) bool { return h[i].c > h[j].c }
func (h charHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *charHeap) Push(x interface{}) {
	*h = append(*h, x.(char))
}

func (h *charHeap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]
	return res
}

func reorganizeString(s string) string {
	cnt, n := make([]int, 26), len(s)
	h, prev, res := &charHeap{}, char{}, make([]rune, 0)
	heap.Init(h)
	for _, ch := range s {
		cnt[ch-'a']++
		if n%2 == 0 && cnt[ch-'a'] > n/2 {
			return ""
		}
		if n%2 == 1 && cnt[ch-'a'] > n/2+1 {
			return ""
		}
	}

	for i := range cnt {
		if cnt[i] != 0 {
			heap.Push(h, char{rune(i + 97), cnt[i]})
		}
	}
	for len(*h) > 0 {
		cur := heap.Pop(h).(char)
		res = append(res, cur.ch)
		cur.c--
		if prev.ch == '0' {
			prev = cur
		} else {
			if prev.c > 0 {
				heap.Push(h, prev)
			}
			prev = cur
		}
	}
	return string(res)
}
