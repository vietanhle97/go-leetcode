package Rotting_Oranges_994

import (
	"container/heap"
	"math"
)

type Rotten struct {
	i   int
	j   int
	m   int
	ind int
}

type Coor struct {
	i int
	j int
}

type RottenQueue []*Rotten

func (pq RottenQueue) Len() int { return len(pq) }

func (pq RottenQueue) Less(i, j int) bool { return pq[i].m < pq[j].m }

func (pq RottenQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].ind = i
	pq[j].ind = j
}

func (pq *RottenQueue) Push(x interface{}) {
	n := len(*pq)
	rotten := x.(*Rotten)
	rotten.ind = n
	*pq = append(*pq, rotten)
}

func (pq *RottenQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	rotten := old[n-1]
	old[n-1] = nil
	rotten.ind = -1
	*pq = old[0 : n-1]
	return rotten
}

func isValid(grid [][]int, i, j, m, n int, check map[Coor]bool) bool {
	if i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == 1 && !check[Coor{i, j}] {
		return true
	}
	return false
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func orangesRotting(grid [][]int) int {
	init := make(RottenQueue, 0)
	check := map[Coor]bool{}
	dir := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	m := len(grid)
	n := len(grid[0])
	heap.Init(&init)
	res := 0
	cnt := 1
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 || grid[i][j] == 2 {
				check[Coor{i, j}] = false
			}
			if grid[i][j] == 2 {
				heap.Push(&init, &Rotten{i, j, 0, cnt})
				cnt += 1
			}
		}
	}

	for init.Len() > 0 {
		r := heap.Pop(&init).(*Rotten)
		if check[Coor{r.i, r.j}] {
			res = min(res, r.m)
		} else {
			check[Coor{r.i, r.j}] = true
			res = max(res, r.m)
		}

		for _, e := range dir {
			if isValid(grid, r.i+e[0], r.j+e[1], m, n, check) {
				heap.Push(&init, &Rotten{r.i + e[0], r.j + e[1], r.m + 1, cnt})
				cnt += 1
			}
		}
	}
	for _, v := range check {
		if !v {
			return -1
		}
	}
	return res
}
