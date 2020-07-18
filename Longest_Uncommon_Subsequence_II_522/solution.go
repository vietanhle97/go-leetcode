package Longest_Uncommon_Subsequence_II_522

import "sort"

func isSub(s1, s2 string) bool {
	j := 0
	for i := range s1 {
		if j < len(s2) && s1[i] == s2[j] {
			j++
		}
	}
	return j == len(s2)
}

func findLUSlength(strs []string) int {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) > len(strs[j])
	})

	for i, s1 := range strs {
		res := true
		for j, s2 := range strs {
			if i != j && len(s2) >= len(s1) && isSub(s2, s1) {
				res = false
				break
			}
		}
		if res {
			return len(s1)
		}
	}

	return -1
}
