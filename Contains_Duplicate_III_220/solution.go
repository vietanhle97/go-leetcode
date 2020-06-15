package Contains_Duplicate_III_220

import "math"

func diff(a, b, t int) bool {
	return int(math.Abs(float64(a)-float64(b))) <= t
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	bucket := map[int]int{}

	for i, e := range nums {
		bucketNum, offSet := e, 0
		if t != 0 {
			bucketNum = e / t
			offSet = 1
		}
		for j := bucketNum - offSet; j < bucketNum+offSet+1; j++ {
			if _, ok := bucket[j]; ok && diff(bucket[j], nums[i], t) {
				return true
			}
		}
		bucket[bucketNum] = nums[i]
		if len(bucket) > k {
			if t == 0 {
				delete(bucket, nums[i-k])
			} else {
				delete(bucket, nums[i-k]/t)
			}
		}
	}
	return false
}
