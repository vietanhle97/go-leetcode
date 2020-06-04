package Longest_Palindromic_Subsequence_516

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestPalindromeSubseq(s string) int {
	n := len(s)

	arr := make([]int, n)

	for i := n - 1; i > -1; i-- {
		tmp := 0
		for j := i; j < n; j++ {
			if j == i {
				arr[j] = 1
			} else if s[i] == s[j] {
				t := arr[j]
				arr[j] = tmp + 2
				tmp = t
			} else {
				tmp = arr[j]
				arr[j] = max(arr[j-1], arr[j])
			}
		}
	}
	return arr[n-1]
}
