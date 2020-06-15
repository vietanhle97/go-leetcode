package Reverse_Substrings_Between_Each_Pair_Of_Parentheses_1190

import (
	"strings"
)

func reverse(l *[]byte, start, end int) {
	for start < end {
		(*l)[start], (*l)[end] = (*l)[end], (*l)[start]
		start++
		end--
	}
}

func reverseParentheses(s string) string {
	l := make([]byte, len(s))
	for i := range s {
		l[i] = s[i]
	}
	stack := make([]int, 0)
	for i, e := range l {
		if e == '(' {
			stack = append(stack, i)
		} else if e == ')' {
			start := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			reverse(&l, start+1, i)
		}
	}
	res := strings.Replace(string(l), "(", "", len(l))
	return strings.Replace(res, ")", "", len(res))

}
