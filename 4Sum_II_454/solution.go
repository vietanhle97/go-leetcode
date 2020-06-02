package _Sum_II_454

func fourSumCount(A []int, B []int, C []int, D []int) int {

	dist := map[int]int{}
	cnt := 0

	for _, a := range A {
		for _, b := range B {
			if _, ok := dist[a+b]; !ok {
				dist[a+b] = 1
			} else {
				dist[a+b] += 1
			}
		}
	}

	for _, c := range C {
		for _, d := range D {
			if _, ok := dist[-c-d]; ok {
				cnt += dist[-c-d]
			}
		}
	}
	return cnt
}
