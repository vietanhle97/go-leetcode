package main

// this file is like a scratch for some techniques I used in the problems above

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const MaxInt = 1 << 31

var k int
var n int
var A []int

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

func main() {
	_, e := fmt.Scanf("%d%d", &n, &k)
	if e != nil {
		return
	}
	if n <= 1 {
		fmt.Println(0)
		return
	}
	A = make([]int, n+1)
	for i := 1; i <= n; i++ {
		_, e := fmt.Scanf("%d", &A[i])
		if e != nil {
			return
		}
	}
	if n == 2 {
		fmt.Println(abs(A[2] - A[1]))
		return
	}
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = MaxInt
	}
	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if i-j >= 1 {
				dp[i] = min(dp[i], dp[i-j]+abs(A[i]-A[i-j]))
			}
		}

	}
	fmt.Println(dp[n])
}
