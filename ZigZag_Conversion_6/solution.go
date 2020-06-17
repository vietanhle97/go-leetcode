package ZigZag_Conversion_6

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 0 || numRows == 1 {
		return s
	}
	l := make([]string, numRows)
	i := 0
	carrier := 0
	dir := 0
	for i < len(s) {
		l[carrier] = l[carrier] + string(s[i])
		if dir == 0 {
			carrier++
		} else {
			carrier--
		}
		if carrier == numRows-1 {
			dir = 1
		}
		if carrier == 0 {
			dir = 0
		}
		i++
	}
	return strings.Join(l, "")
}
