package Basic_Calculator_II_227

import (
	"strconv"
	"strings"
)

type token struct {
	num  int
	val  string
	ind  int
	Type string
}

func (t *token) isNumber() bool {
	return t.Type == "number"
}

func isDigit(ch string) bool {
	return strings.Contains("0123456789", ch)
}

func findNextToken(s string, pos int) (token, int) {
	for pos < len(s) && s[pos] == ' ' {
		pos++
	}
	if pos == len(s) {
		return token{-1, "", pos, "stop"}, pos
	}
	begin := pos
	ch := string(s[pos])
	if isDigit(ch) {
		for pos < len(s) && isDigit(string(s[pos])) {
			pos++
		}
		val, _ := strconv.Atoi(s[begin:pos])
		return token{val, "", begin, "number"}, pos
	}
	return token{-1, string(s[pos]), begin, "symbol"}, pos + 1

}

func tokenize(s string) []token {
	tokens := make([]token, 0)
	pos := 0
	for {
		tok, nPos := findNextToken(s, pos)
		tokens = append(tokens, tok)
		if tok.Type == "stop" {
			return tokens
		}
		pos = nPos
	}
}

func parseItem(tokens *[]token) int {
	t := (*tokens)[0]
	*tokens = (*tokens)[1:]
	if t.isNumber() {
		return t.num
	}

	expr := parseExpression(tokens)
	*tokens = (*tokens)[1:]
	return expr
}

func parseTerm(tokens *[]token) int {
	t := (*tokens)[0]
	sign := 1
	if t.val == "-" {
		sign = -1
	}
	if t.val == "+" || sign < 0 {
		*tokens = (*tokens)[1:]
	}
	res := parseItem(tokens) * sign
	for (*tokens)[0].val == "*" || (*tokens)[0].val == "/" {
		op := (*tokens)[0].val
		*tokens = (*tokens)[1:]
		rhs := parseItem(tokens)
		if op == "*" {
			res = res * rhs
		} else {
			res = res / rhs
		}
	}
	return res
}

func parseExpression(tokens *[]token) int {
	res := parseTerm(tokens)
	t := (*tokens)[0]
	for t.val == "+" || t.val == "-" {
		*tokens = (*tokens)[1:]
		rhs := parseTerm(tokens)
		if t.val == "+" {
			res += rhs
		} else {
			res -= rhs
		}
		t = (*tokens)[0]
	}
	return res
}

func calculate(s string) int {
	tokens := tokenize(s)
	return parseExpression(&tokens)
}
