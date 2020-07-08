package Length_of_Longest_Fibonacci_Subsequence_873

func FibBackTrack(start int, max_ int, nums []int, path []int, res *int) {
	if len(path) > *res {
		*res = len(path)
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
	res := 0
	FibBackTrack(0, len(A), A, []int{}, &res)
	if res > 2 {
		return res
	}
	return 0

}
