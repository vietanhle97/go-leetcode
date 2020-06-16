package Reverse_Words_In_A_String_151

import "strings"

func reverseWords(s string) string {
	l := strings.Split(s, " ")
	i := len(l) - 1
	res := ""
	b := false
	for i >= 0 {
		if len(l[i]) != 0 {
			if b {
				res += " "
			}
			res += l[i]
			b = true
		}
		i--
	}
	return res
}
