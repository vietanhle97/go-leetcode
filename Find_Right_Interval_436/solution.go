package Find_Right_Interval_436

import (
	"sort"
)

type mem struct {
	x int
	y int
}

type Mem []mem

func (a Mem) Len() int { return len(a) }
func (a Mem) Less(i, j int) bool {
	return a[i].x < a[j].x
}
func (a Mem) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func binarySearch(Mem []mem, m mem, l, r int) int {
	res := -1
	for l <= r {
		mid := l + (r-l)/2
		if Mem[mid].x == m.y {
			return mid
		} else if Mem[mid].x > m.y {
			if res == -1 {
				res = mid
			}
			if Mem[mid].x < Mem[res].x {
				res = mid
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}

func findRightInterval(intervals [][]int) []int {
	if len(intervals) == 1 {
		return []int{-1}
	}
	m := len(intervals)
	ind := map[mem]int{}
	M := make([]mem, m)
	for i, e := range intervals {
		tmp := mem{e[0], e[1]}
		M[i] = tmp
		ind[tmp] = i
	}
	sort.Sort(Mem(M))
	res := make([]int, m)
	for i, e := range M {
		tmp := binarySearch(M, e, i, m-1)
		if tmp != -1 {
			res[ind[e]] = ind[M[tmp]]
		} else {
			res[ind[e]] = -1
		}
	}

	return res
}
