package Binary_Trees_With_Factors_823

import (
	"math"
	"sort"
)

func numFactoredBinaryTrees(A []int) int {
	sort.Ints(A)
	table := make([]int64, len(A))
	for i := range A {
		table[i] = 1
	}
	m := map[int]int{}
	res := int64(0)
	for i := range A {
		m[A[i]] = i
	}
	for i := 1; i < len(A); i++ {
		for j := 0; j < i; j++ {
			if A[i]%A[j] == 0 {
				if _, ok := m[A[i]/A[j]]; ok {
					table[i] += table[j] * table[m[A[i]/A[j]]]
				}
			}
		}
	}
	for _, e := range table {
		res += e
	}
	return int(res % int64(math.Pow10(9)+7))
}
