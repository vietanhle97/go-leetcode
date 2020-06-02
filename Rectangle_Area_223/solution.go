package Rectangle_Area_223

func min(A, B int) int {
	if A < B {
		return A
	}
	return B
}
func max(A, B int) int {
	if A > B {
		return A
	}
	return B
}
func overlap(A, B, C, D, E, F, G, H int) bool {
	overlapX := (E >= A && E <= C) || (E <= A && G >= C) || (G >= A && G <= C)
	overlapY := (F >= B && F <= D) || (F <= B && H >= D) || (H >= B && H <= D)
	return overlapX && overlapY
}

func findOverlappedRange(A, B, C, D, E, F, G, H int) int {
	minX := max(A, E)
	maxX := min(C, G)
	minY := max(B, F)
	maxY := min(D, H)
	return (maxX - minX) * (maxY - minY)
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	firstRecArea := (C - A) * (D - B)
	secondRecArea := (G - E) * (H - F)
	total := firstRecArea + secondRecArea
	isOverlapped := overlap(A, B, C, D, E, F, G, H)
	if !isOverlapped {
		return total
	}
	return total - findOverlappedRange(A, B, C, D, E, F, G, H)

}
