package main

// this file is like a scratch for some techniques I used in the problems above

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const MaxInt = 1 << 31

var k, N, M, n, u, v int
var A, B int
var T int64
var m int64

type num struct {
	s string
}

type Pair struct {
	prev int
	val  int
	cnt  int
}

type maxHeap []int

func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	n := h.Len()
	res := (*h)[n-1]
	*h = (*h)[:n-1]
	return res
}

type Nums []num

func (a Nums) Len() int { return len(a) }
func (a Nums) Less(i, j int) bool {
	n1, _ := strconv.Atoi(a[i].s + a[j].s)
	n2, _ := strconv.Atoi(a[j].s + a[i].s)
	return n1 < n2
}
func (a Nums) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func test(nums []int) {
	m := len(nums)
	N := make([]num, 0)
	for i := range nums {
		N = append(N, num{strconv.Itoa(nums[i])})
	}
	sort.Sort(Nums(N))
	res := ""
	for i := m - 1; i > -1; i-- {
		if i == 0 {
			res += N[i].s
			break
		}
		if len(res) == 0 && N[i].s == "0" {
			continue
		}
		res += N[i].s
	}
	fmt.Println(res)
}

func testLengthOfLongestSubstring(s string) string {
	//start := time.Now()
	m := len(s)
	if m < 2 {
		return s
	}

	table := make([][]int, m)
	max_ := 0
	st, en := 0, 0
	for i := 0; i < m; i++ {
		table[i] = make([]int, m)
		for j := i; j < m; j++ {
			if j == i {
				table[i][j] = 1
			} else {
				if !strings.Contains(s[i:j], string(s[j])) {
					table[i][j] = table[i][j-1] + 1
				} else {
					table[i][j] = 1
				}
				if table[i][j] > max_ {
					st = i
					en = j
					max_ = table[i][j]
				}
			}
		}
	}
	//end := time.Now()
	return s[st : en+1]
}

func countCombination(MAX, n int) int {
	res := 1
	i := 1
	for i <= n {
		res *= MAX
		i++
		MAX--
	}
	return res
}

func getCombination(MAX, n, k int) []int {
	nums := make([]int, 0)
	for i := 0; i < MAX; i++ {
		nums = append(nums, i+1)
	}
	res := make([]int, 0)
	if k > countCombination(MAX, n) || k <= 0 {
		return res
	}
	i := MAX
	j := n
	for len(res) < n {
		fac := countCombination(i-1, j-1)
		i--
		j--
		ind := k / fac
		r := k % fac
		if ind > 0 && r == 0 {
			ind -= 1
		}
		fmt.Println(fac, nums, ind, res)
		res = append(res, nums[ind])
		k -= fac * ind

		nums = append(nums[:ind], nums[ind+1:]...)
	}
	return res
}

func factorial(n int) int {
	res := 1
	for n > 0 {
		res *= n
		n--
	}
	return res
}

func getPermutation(n int, k int) string {
	nums := make([]int, n)
	for i, _ := range nums {
		nums[i] = i + 1
	}
	res := ""
	for len(res) < n {
		fac := factorial(len(nums) - 1)
		ind := k / fac
		r := k % fac
		if ind > 0 && r == 0 {
			ind -= 1
		}
		res += strconv.Itoa(nums[ind])
		k -= fac * ind

		nums = append(nums[:ind], nums[ind+1:]...)
	}
	return res
}

func convertPoint(n, i, j int) int {
	return n*i + j
}

func isValid(m, n, i, j int) bool {
	return i >= 0 && i < m && j >= 0 && j < n
}

func dfsCycle(p int, visited, rec []bool, graph map[int][]int) bool {
	visited[p], rec[p] = true, true
	for _, e := range graph[p] {
		if visited[e] == false {
			if dfsCycle(e, visited, rec, graph) {
				return true
			}
		}
		if rec[e] {
			return true
		}
	}
	rec[p] = false
	return false
}

func dfs(u int, vis, dist []int, A [][]int) {
	vis[u] = 1
	for _, v := range A[u] {
		if vis[v] == 0 {
			dist[v] = dist[u] + 1
			dfs(v, vis, dist, A)
		}
	}
}
func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func bs(b []int64, v int64) int {
	l, r := 0, len(b)-1
	for l < r {
		mid := l + (r-l)/2
		if b[mid] == v {
			return mid
		} else if b[mid] < v {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func sumMin(a []int) int64 {
	n := len(a)
	l, r, q := make([]int, n), make([]int, n), make([]int, 0)
	for i := 0; i < n; i++ {
		for len(q) > 0 && a[q[len(q)-1]] >= a[i] {
			l[i] += l[q[len(q)-1]] + 1
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	q = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(q) > 0 && a[q[len(q)-1]] > a[i] {
			r[i] += r[q[len(q)-1]] + 1
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	res := int64(0)
	for i := 0; i < n; i++ {
		res += int64(l[i]+1) * int64(r[i]+1) * int64(a[i])
	}
	return res
}
func sumMax(a []int) int64 {
	n := len(a)
	l, r, q := make([]int, n), make([]int, n), make([]int, 0)
	for i := 0; i < n; i++ {
		for len(q) > 0 && a[q[len(q)-1]] <= a[i] {
			l[i] += l[q[len(q)-1]] + 1
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	q = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(q) > 0 && a[q[len(q)-1]] < a[i] {
			r[i] += r[q[len(q)-1]] + 1
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	res := int64(0)
	for i := 0; i < n; i++ {
		res += int64(l[i]+1) * int64(r[i]+1) * int64(a[i])
	}
	return res
}

func calculateMaxMinusMinOfAllSubsequence2(a []int) int64 {
	MIN, MAX := sumMin(a), sumMax(a)
	return MAX - MIN
}

func main() {
	a := []int{2, 2, 3, 5}
	res := calculateMaxMinusMinOfAllSubsequence2(a)
	fmt.Println(res)
}
