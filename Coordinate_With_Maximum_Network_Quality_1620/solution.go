package Coordinate_With_Maximum_Network_Quality_1620

import "math"

func dist(a, b []int) float64 {
	return math.Sqrt(float64((a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])))
}

func comp(a, b []int) []int {
	if len(a) == 0 {
		return b
	}
	if a[0] < b[0] {
		return a
	}
	if a[0] == b[0] && a[1] < b[1] {
		return a
	}
	return b
}

func bestCoordinate(towers [][]int, radius int) []int {
	res := make([]int, 0)
	max := -(1<<31 - 1) - 1
	for i := range towers {
		a := []int{towers[i][0], towers[i][1]}
		tmp := 0
		for _, b := range towers {
			d := dist(a, b)
			if d <= float64(radius) {
				tmp += int(float64(b[2]) / (1 + dist(a, b)))
			}
		}
		if tmp >= max {
			if tmp == max {
				res = comp(res, a)
			} else {
				res = a
			}

			max = tmp
		}
	}
	return res
}
