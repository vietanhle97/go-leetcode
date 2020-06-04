package _Sum_With_Multiplicity_923

import "math"

func threeSumMulti(A []int, target int) int {
	d := [101]int{}
	res := 0
	MOD := int(math.Pow10(9)) + 7
	for _, e := range A {
		d[e] += 1
	}
	//x != y != z
	for x := 0; x < 101; x++ {
		for y := x + 1; y < 101; y++ {
			z := target - x - y
			if y >= 0 && y < z && z <= 100 {
				res += d[x] * d[y] * d[z]
				res %= MOD
			}
		}
	}

	// x != y == z
	for x := 0; x < 101; x++ {
		if (target-x)%2 == 0 {
			y := (target - x) / 2
			if x < y && y <= 100 && y >= 0 {
				res += d[x] * d[y] * (d[y] - 1) / 2
				res %= MOD
			}
		}
	}

	// x == y
	for x := 0; x < 101; x++ {
		z := target - 2*x
		if z <= 100 && z >= 0 && x < z {
			res += d[x] * (d[x] - 1) * d[z] / 2
			res %= MOD
		}
	}

	//x == y == z
	if target%3 == 0 {
		x := target / 3
		if x >= 0 && x <= 100 {
			res += d[x] * (d[x] - 1) * (d[x] - 2) / 6
			res %= MOD
		}
	}
	return res
}
