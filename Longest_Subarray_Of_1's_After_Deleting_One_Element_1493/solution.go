package Longest_Subarray_Of_1_s_After_Deleting_One_Element_1493

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func longestSubarray(nums []int) int {
	cnt := make([]int, 0)
	m := len(nums)
	i := 0
	res := 0
	for i < m {
		if nums[i] == 0 {
			cnt = append(cnt, -1)
			i++
		} else {
			c := 0
			for nums[i] == 1 {
				c, i = c+1, i+1
				if i >= m {
					break
				}
			}
			cnt = append(cnt, c)
		}
	}
	n := len(cnt)
	if n < 3 {
		for _, e := range cnt {
			if e != -1 {
				if n == 2 {
					return e
				} else {
					return e - 1
				}
			}
		}
		return 0
	}

	if cnt[0] != -1 {
		res = max(res, cnt[0])
	}
	if cnt[n-1] != -1 {
		res = max(res, cnt[n-1])
	}
	for i := 1; i < n-1; i++ {
		if cnt[i] == -1 {
			if cnt[i-1] != -1 && cnt[i+1] != -1 {
				res = max(res, cnt[i-1]+cnt[i+1])
			} else {
				continue
			}
		} else {
			res = max(res, cnt[i])
		}
	}
	return res
}
