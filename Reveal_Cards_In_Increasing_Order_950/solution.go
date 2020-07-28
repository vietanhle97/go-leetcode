package Reveal_Cards_In_Increasing_Order_950

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	N := len(deck)
	ind := make([]int, N)
	for i := range ind {
		ind[i] = i
	}
	ans := make([]int, N)
	sort.Ints(deck)

	for _, e := range deck {
		ans[ind[0]] = e
		ind = ind[1:]
		if len(ind) > 0 {
			tmp := ind[0]
			ind = ind[1:]
			ind = append(ind, tmp)
		}
	}
	return ans
}
