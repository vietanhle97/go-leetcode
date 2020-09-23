package Maximum_XOR_Of_Two_Numbers_In_An_Array_421

func findMaximumXOR(nums []int) int {
	max, mask := 0, 0
	check := map[int]bool{}

	for i := 30; i >= 0; i-- {
		mask |= 1 << i
		nMax := max | (1 << i)
		for _, e := range nums {
			check[e&mask] = true
		}
		for k, _ := range check {
			if _, ok := check[nMax^k]; ok {
				max = nMax
				break
			}
		}
		check = map[int]bool{}
	}
	return max
}
