package Restore_IP_Addresses_93

import (
	"strconv"
)

func backtrack(start, cnt int, s string, path string, res *[]string) {
	if cnt > 3 {
		return
	}
	if start >= len(s) && cnt == 3 && path[len(path)-1] != '.' {
		*res = append(*res, path)
		return
	}
	for i := start; i < len(s); i++ {
		cur := s[start : i+1]
		if len(cur) == 1 {
			if cnt < 3 {
				backtrack(i+1, cnt+1, s, path+cur+".", res)
			} else if cnt == 3 && i == len(s)-1 {
				backtrack(i+1, cnt, s, path+cur, res)
			}

		}
		if len(cur) > 1 && cur[0] != '0' {
			v, err := strconv.ParseInt(cur, 10, 64)
			if err == nil && v <= 255 && v >= 0 {
				if cnt < 3 {
					backtrack(i+1, cnt+1, s, path+cur+".", res)
				} else if cnt == 3 && i == len(s)-1 {
					backtrack(i+1, cnt, s, path+cur, res)
				}
			}
		}
	}
}

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	backtrack(0, 0, s, "", &res)
	return res
}
