package Evaluate_Reverse_Polish_Notation_150

import (
	"strconv"
)

func isNum(n string) (bool, int) {
	if val, err := strconv.Atoi(n); err == nil {
		return true, val
	} else {
		return false, 0
	}
}

func eval(s1, s2 int, s3 string) int {
	if s3 == "+" {
		return s1 + s2
	} else if s3 == "*" {
		return s1 * s2
	} else if s3 == "-" {
		return s1 - s2
	} else {
		return s1 / s2
	}
}

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	res := 0
	for _, e := range tokens {
		b, v := isNum(e)
		if b {
			stack = append(stack, v)
		} else {
			first := stack[len(stack)-1]
			second := stack[len(stack)-2]
			res = eval(second, first, e)
			stack = append(stack[:len(stack)-2], res)
		}
	}
	if len(stack) > 0 {
		return stack[len(stack)-1]
	}
	return res
}
