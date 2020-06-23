package Median_Of_Two_Sorted_Arrays_4

import (
	"math"
)

const UintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64

const (
	MaxInt = 1<<(UintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
	MinInt = -MaxInt - 1         // -1 << 31 or -1 << 63
)

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}

	k := (m + n - 1) / 2
	l := 0
	r := min(k, m)

	for l < r {
		mid1 := l + (r-l)/2
		mid2 := k - mid1
		if nums1[mid1] > nums2[mid2] {
			r = mid1
		} else {
			l = mid1 + 1
		}
	}
	tmpMin1 := MinInt
	tmpMin2 := MinInt
	if l >= 1 {
		tmpMin1 = nums1[l-1]
	}
	if k >= l {
		tmpMin2 = nums2[k-l]
	}
	a := max(tmpMin1, tmpMin2)
	if (n+m)%2 != 0 {
		return float64(a)
	}
	tmpMax1 := MaxInt
	tmpMax2 := MaxInt
	if l < m {
		tmpMax1 = nums1[l]
	}
	if k-l+1 < n {
		tmpMax2 = nums2[k-l+1]
	}
	b := min(tmpMax1, tmpMax2)
	return float64(a+b) / 2
}
