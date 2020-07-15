package Best_Position_For_A_Service_Centre_1515

import "math"

func dist(x, y float64, p []float64) float64 {
	distX := math.Abs(x - p[0])
	distY := math.Abs(y - p[1])
	return math.Sqrt(distX*distX + distY*distY)
}

func getMinDistSum(positions [][]int) float64 {
	n := len(positions)
	X, Y := 0.0, 0.0
	for _, e := range positions {
		X += float64(e[0])
		Y += float64(e[1])
	}
	X /= float64(n)
	Y /= float64(n)

	pX, pY := math.Inf(-1), math.Inf(-1)
	limit := 1e-8
	for dist(X, Y, []float64{pX, pY}) > limit {
		upX, upY, down := 0.0, 0.0, 0.0
		for _, e := range positions {
			p := []float64{float64(e[0]), float64(e[1])}
			tmp := dist(X, Y, p)
			if tmp == 0 {
				break
			}
			upX += p[0] / tmp
			upY += p[1] / tmp
			down += 1 / tmp
		}
		if down == 0 {
			break
		}
		nX, nY := upX/down, upY/down
		pX, pY, X, Y = X, Y, nX, nY
	}
	min := 0.0
	for _, e := range positions {
		p := []float64{float64(e[0]), float64(e[1])}
		min += dist(X, Y, p)
	}
	return min
}
