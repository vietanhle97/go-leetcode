package Array_Of_Doubled_Pairs_954

import (
	"math"
)

func partition(nums []int, l, h int) int {
	i := l - 1
	pivot := nums[h]
	for j := l; j < h; j++ {
		if int(math.Abs(float64(nums[j]))) > int(math.Abs(float64(pivot))) {
			i += 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[h] = nums[h], nums[i+1]
	return i + 1
}

func quicksort(nums []int, l, h int) {
	if l < h {
		pivot := partition(nums, l, h)
		quicksort(nums, l, pivot-1)
		quicksort(nums, pivot+1, h)
	}
}

func canReorderDoubled(A []int) bool {
	quicksort(A, 0, len(A)-1)
	odd := map[int]int{}
	even := map[int]int{}

	for _, e := range A {
		if e%2 == 1 {
			if _, ok := even[2*e]; ok && even[2*e] > 0 {
				even[2*e] -= 1
			} else {
				if _, ok := odd[e]; !ok {
					odd[e] = 1
				} else {
					odd[e] += 1
				}
			}

		} else {
			if _, ok := odd[e/2]; ok && odd[e/2] > 0 {
				odd[e/2] -= 1
			} else if _, ok := even[e/2]; ok && even[e/2] > 0 {
				even[e/2] -= 1
			} else if _, ok := even[2*e]; ok && even[2*e] > 0 {
				even[2*e] -= 1
			} else {
				if _, ok := even[e]; !ok {
					even[e] = 1
				} else {
					even[e] += 1
				}
			}
		}
	}
	for _, v := range odd {
		if v > 0 {
			return false
		}
	}
	for _, v := range even {
		if v > 0 {
			return false
		}
	}
	return true
}
