package Path_In_Zigzag_Labelled_Binary_Tree_1104

import "math"

func pow2(a int) int {
	return int(math.Pow(2, float64(a)))
}

func reverse(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}

func pathInZigZagTree(label int) []int {
	res := []int{label}
	if label == 1 {
		return res
	}
	i := 1
	cnt := -1
	for i <= label {
		i *= 2
		cnt++
	}
	carry := 1
	for cnt > 0 {
		label /= 2
		if carry == 1 {
			l := pow2(cnt - 1)
			r := 2*l - 1
			if label == r+1 {
				res = append(res, 2*r+1)
			} else {
				res = append(res, l+r-label)
			}
			carry = 0
		} else {
			res = append(res, label)
			carry = 1
		}
		cnt--
	}
	return reverse(res)
}
