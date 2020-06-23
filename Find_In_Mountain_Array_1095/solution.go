package Find_In_Mountain_Array_1095

import "math"

type MountainArray struct {
	nums []int
}

func (this *MountainArray) get(index int) int { return this.nums[index] }
func (this *MountainArray) length() int       { return len(this.nums) }

// Solution starts from here
// I wrote nums inside MountainArray and the return statements to their function to make my tests runnable

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func findPivotPoint(mountainArr *MountainArray, l, r, target int) (int, int) {
	first := -1
	for l < r {
		mid := l + (r-l)/2
		tmp := mountainArr.get(mid)
		if tmp == target {
			if first == -1 {
				first = mid
			}
			first = min(first, mid)
		}
		if tmp < mountainArr.get(mid+1) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l, first
}

func binarySearchInc(mountainArr *MountainArray, l, r, target int) int {
	for l <= r {
		mid := l + (r-l)/2
		tmp := mountainArr.get(mid)
		if tmp == target {
			return mid
		} else if tmp > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
func binarySearchDec(mountainArr *MountainArray, l, r, target int) int {
	for l <= r {
		mid := l + (r-l)/2
		tmp := mountainArr.get(mid)
		if tmp == target {
			return mid
		} else if tmp > target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	m := mountainArr.length()
	pivot, first := findPivotPoint(mountainArr, 0, mountainArr.length()-1, target)
	if first == -1 {
		ls := binarySearchInc(mountainArr, 0, pivot-1, target)
		if ls != -1 {
			return ls
		}
		return binarySearchDec(mountainArr, pivot+1, m-1, target)
	}
	if first <= pivot {
		return first
	} else {
		ls := binarySearchInc(mountainArr, 0, pivot-1, target)
		if ls != -1 {
			return ls
		}
		return first
	}
}
