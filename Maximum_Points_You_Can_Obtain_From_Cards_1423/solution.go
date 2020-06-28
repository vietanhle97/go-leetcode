package Maximum_Points_You_Can_Obtain_From_Cards_1423

import "math"

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func findSum(cardPoints []int) int {
	sum := 0
	for _, e := range cardPoints {
		sum += e
	}
	return sum
}

func maxScore(cardPoints []int, k int) int {
	m := len(cardPoints)
	maxPoints := findSum(cardPoints)
	if k == m {
		return maxPoints
	}
	res := findSum(cardPoints[0:(m - k)])
	cur := res
	for i := m - k; i < m; i++ {
		res = min(res, cur+cardPoints[i]-cardPoints[i-(m-k)])
		cur = cur + cardPoints[i] - cardPoints[i-(m-k)]
	}
	return maxPoints - res
}
