package Next_Greater_Element_III_556

import (
	"math"
	"sort"
	"strconv"
)

func nextGreaterElement(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	l := make([]int, m)
	for i := range s {
		l[i], _ = strconv.Atoi(string(s[i]))
	}
	first := -1
	for i := len(l) - 1; i > 0; i-- {
		if l[i-1] < l[i] {
			first = i - 1
			break
		}
	}
	if first == -1 {
		return -1
	}
	min := -1
	for i := first; i < len(l); i++ {
		if l[i] > l[first] {
			min = i
		}
	}

	l[first], l[min] = l[min], l[first]
	tmp := make([]int, m-first-1)
	copy(tmp, l[first+1:])
	sort.Ints(tmp)
	l = append(l[:first+1], tmp...)
	res := 0
	for i := range l {
		res += int(math.Pow10(m-i-1)) * l[i]
	}
	if res < int(math.Pow(2, 31))-1 {
		return res
	}
	return -1
}
