package Pow_x__n__50

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	tmp := myPow(x, n/2)
	if n%2 == 0 {
		return tmp * tmp
	} else {
		if n > 0 {
			return tmp * tmp * x
		} else {
			return tmp * tmp / x
		}
	}
}
