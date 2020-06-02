package Increasing_Triplet_Subsequence_334

func increasingTriplet(nums []int) bool {
	first, second := int(^uint(0)>>1), int(^uint(0)>>1)

	for _, e := range nums {
		if e <= first {
			first = e
		} else if e <= second {
			second = e
		} else {
			return true
		}
	}
	return false
}
