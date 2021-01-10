package Count_Good_Meals_1711

const mod = int(1e9 + 7)

func countPairs(D []int) int {
	pow := []int{1}
	for i := 0; i < 21; i++ {
		pow = append(pow, pow[len(pow)-1]*2)
	}
	res, m := 0, map[int]int{}

	for _, e := range D {
		for _, p := range pow {
			if p >= e {
				if _, ok := m[p-e]; ok {
					res += m[p-e]
					res %= mod
				}
			}
		}
		m[e]++
	}
	return res
}
