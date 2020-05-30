package main

import "fmt"

type Res struct {
	val int
}

func FibBackTrack(start int, max_ int, nums []int, path []int, res *Res) {
	if len(path) > res.val {
		res.val = len(path)
	}
	if start >= max_ {
		return
	}
	for i := start; i < max_; i++ {
		if len(path) < 2 {
			FibBackTrack(i+1, max_, nums, append(path, nums[i]), res)
		} else {
			if nums[i] == path[len(path)-1]+path[len(path)-2] {
				FibBackTrack(i+1, max_, nums, append(path, nums[i]), res)
			}
		}
	}

}
func lenLongestFibSubseq(A []int) int {
	res := Res{0}
	FibBackTrack(0, len(A), A, []int{}, &res)
	if res.val > 2 {
		return res.val
	}
	return 0

}

func main() {
	A := []int{12, 13, 15, 18, 26, 31, 34, 35, 37, 39, 41, 42, 45, 47, 49, 50, 52, 56, 57, 60, 61, 63, 69, 72, 78, 81, 83, 91, 97, 102, 107, 110, 113, 117, 126, 128, 140, 143, 153, 162, 171, 176, 179, 195, 209, 223, 234, 250, 264, 281, 286, 289, 312, 321, 337, 363, 377, 403, 426, 452, 465, 507, 516, 546, 586, 653, 690, 733, 751, 754, 819, 837, 883, 949, 1056, 1116, 1185, 1216, 1219, 1326, 1353, 1429, 1535, 1709, 1806, 1967, 1973, 2145, 2190, 2312, 2922, 3183, 3192, 3471, 3543, 3741, 4728, 5150, 5165, 5616, 5733, 6053, 7650, 8333, 8357, 9087, 9276, 13483, 13522, 14703, 15009, 21816, 21879, 23790, 24285, 35299, 35401, 38493, 39294, 63579}
	fmt.Println(lenLongestFibSubseq(A))
}
