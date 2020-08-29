package Pancake_Sorting_969

func flip(A []int, j int) {
	l, r := 0, j
	for l < r {
		A[l], A[r] = A[r], A[l]
		l++
		r--
	}
}

func findMax(A []int, j int) int {
	res, ind := A[0], 0
	for i := 0; i < j; i++ {
		if A[i] > res {
			res = A[i]
			ind = i
		}
	}
	return ind
}

func pancakeSort(A []int) []int {
	n := len(A)
	res, cnt := make([]int, 0), n
	for i := 0; i < n; i++ {
		tmp := findMax(A, cnt)
		if tmp == cnt-1 {
			cnt--
			continue
		} else {
			flip(A, tmp)
			res = append(res, tmp+1)
			flip(A, cnt-1)
			res = append(res, cnt)
		}
		cnt--
	}
	return res
}
