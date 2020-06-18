package main

// this file is like a scratch for some techniques I used in the problems above

import (
	"fmt"
	"sort"
	"strconv"
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
	fmt.Println(m, n, p)
}
