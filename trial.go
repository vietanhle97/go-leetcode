package main

// this file is like a scratch for some techniques I used in the problems above

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

type num struct {
	s string
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
	return int(math.Max(float64(a), float64(b)))
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
	nums := []int{824, 938, 1399, 5607, 6973, 5703, 9609, 4398, 8247}
	nums1 := []int{0, 0}
	nums2 := []int{121, 12}
	nums3 := []int{3, 30, 34, 5, 9}
	nums4 := []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	test(nums)
	test(nums1)
	test(nums2)
	test(nums3)
	test(nums4)

	n := []int{1, 2, 32}
	m := make([]int, len(n))
	p := n
	copy(m, n)
	m = append(m, 9)
	p = append(p, 9)
	//fmt.Println(m, n, p)
	//fmt.Println(testLengthOfLongestSubstring(s))
	fmt.Println(countCombination(6, 3))
	fmt.Println(getCombination(3, 2, 1))
	start := time.Now()
	for i := 1; i <= 100; i++ {
		getPermutation(9, i)
	}
	fmt.Println(time.Now().Sub(start))
}
