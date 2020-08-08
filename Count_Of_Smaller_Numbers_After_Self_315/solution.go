package Count_Of_Smaller_Numbers_After_Self_315

const (
	MAX = 1<<31 - 1
	MIN = -MAX - 1
)

func findMin(A []int) int {
	res := MAX
	for _, e := range A {
		if e < res {
			res = e
		}
	}
	return res
}

func count(BIT []int, i int) int {
	i++
	s := 0
	for i > 0 {
		s += BIT[i]
		i -= i & (-i)
	}
	return s
}

func update(BIT []int, n, i int) {
	i++
	for i < n {
		BIT[i]++
		i += i & (-i)
	}
}

func countSmaller(A []int) []int {
	m := len(A)
	if m == 0 {
		return []int{}
	}
	max, min := MIN, findMin(A)
	for i := range A {
		A[i] = A[i] - min + 1
		if A[i] > max {
			max = A[i]
		}
	}

	BIT, res := make([]int, max+1), make([]int, 0)
	for i := m - 1; i >= 0; i-- {
		res = append([]int{count(BIT, A[i]-2)}, res...)
		update(BIT, max+1, A[i]-1)
	}
	return res
}
