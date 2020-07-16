package Multiply_Strings_43

func addString(a, b string) string {
	res := ""
	carrier := 0
	for len(b) < len(a) {
		b = "0" + b
	}
	m := len(a)
	i := m - 1
	for i >= 0 {
		n1, n2 := int(a[i]-'0'), int(b[i]-'0')
		tmp := n1 + n2 + carrier
		if tmp >= 10 {
			carrier = 1
		} else {
			carrier = 0
		}
		res = string(tmp%10+'0') + res
		i--
	}
	if carrier == 1 {
		return "1" + res
	}
	return res
}

func multiString(s1 string, s2 byte) string {
	n := int(s2 - '0')
	if n == 0 {
		return "0"
	}
	res := ""
	for i := 1; i < n+1; i++ {
		if len(res) == 0 {
			res = addString(s1, res)
		} else {
			res = addString(res, s1)
		}
	}
	return res
}
func multiply(num1 string, num2 string) string {
	m := len(num2)
	res := ""
	for i := range num2 {
		tmp := multiString(num1, num2[i])
		for j := 0; j < m-1-i; j++ {
			tmp += "0"
		}
		if len(res) == 0 {
			res = addString(tmp, res)
		} else {
			res = addString(res, tmp)
		}

	}
	for len(res) > 1 {
		if res[0] == '0' {
			res = res[1:]
		} else {
			break
		}
	}
	return res
}
