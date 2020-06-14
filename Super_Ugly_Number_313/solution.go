package Super_Ugly_Number_313

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
func nthSuperUglyNumber(n int, primes []int) int {
	mul := make([]int, len(primes))
	copy(mul, primes)
	h := &IntHeap{}
	check := map[int]bool{}
	for _, e := range mul {
		heap.Push(h, e)
		check[e] = true
	}
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
