package Broken_Calculator_991

func find(X, Y int) int {
	if X >= Y {
		return X - Y
	}
	if Y%2 == 0 {
		return 1 + find(X, Y/2)
	} else {
		return 1 + find(X, Y+1)
	}
}

func brokenCalc(X int, Y int) int {
	if X >= Y {
		return X - Y
	}
	return find(X, Y)
}
