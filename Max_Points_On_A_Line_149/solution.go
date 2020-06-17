package Max_Points_On_A_Line_149

type Coefficient struct {
	a int
	b int
	c int
	x int
	y int
}

func makeCoefficient(x1, y1, x2, y2 int) Coefficient {
	a := y1 - y2
	b := x1*y2 - x2*y1
	c := x1 - x2
	return Coefficient{a, b, c, x1, y1}
}

func maxPoints(points [][]int) int {
	m := len(points)
	if m < 3 {
		return m
	}
	max := 0
	for i := 0; i < m; i++ {
		lineMap := map[Coefficient]int{}
		x1, y1 := points[i][0], points[i][1]
		for j := 0; j < m; j++ {
			if j != i {
				x2, y2 := points[j][0], points[j][1]
				if len(lineMap) == 0 {
					lineMap[makeCoefficient(x1, y1, x2, y2)] = 2
				} else {
					check := false
					for k, _ := range lineMap {
						if k.a == 0 && k.b == 0 && k.c == 0 {
							if x1 == x2 && y1 == y2 {
								lineMap[k] += 1
								check = true
							}
						} else if k.c == 0 {
							if x2 == k.x {
								lineMap[k] += 1
								check = true
							}
						} else if k.a == 0 {
							if y2 == k.y {
								lineMap[k] += 1
								check = true
							}
						} else {
							if k.a*x2+k.b == y2*k.c {
								lineMap[k] += 1
								check = true
							}
						}
					}
					if !check {
						lineMap[makeCoefficient(x1, y1, x2, y2)] = 2
					}
				}
			}
		}
		for _, v := range lineMap {
			if v > max {
				max = v
			}
		}
	}
	return max
}
