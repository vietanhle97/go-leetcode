package Maximum_Area_Of_A_Piece_Of_Cake_After_Horizontal_And_Vertical_Cuts_1465

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxArea(h int, w int, hCuts []int, wCuts []int) int {
	MOD := int(1e9) + 7
	hCuts = append([]int{0}, append(hCuts, h)...)
	wCuts = append([]int{0}, append(wCuts, w)...)
	sort.Ints(hCuts)
	sort.Ints(wCuts)
	maxH, maxW := 0, 0
	for i := 1; i < len(hCuts); i++ {
		hCuts[i-1] = hCuts[i] - hCuts[i-1]
		maxH = max(hCuts[i-1], maxH)
	}
	for i := 1; i < len(wCuts); i++ {
		wCuts[i-1] = wCuts[i] - wCuts[i-1]
		maxW = max(wCuts[i-1], maxW)
	}
	return ((maxH % MOD) * (maxW % MOD)) % MOD
}
