package Maximum_Binary_String_After_Change_1702

func maximumBinaryString(b string) string {
	m := len(b)
	st := make([]byte, m+1)
	prev := -1
	for i := range b {
		if b[i] == '0' {
			if prev == -1 {
				prev = i
			} else {
				prev++
			}
		}
	}
	for i := range st {
		if i != prev {
			st[i] = '1'
		} else {
			st[i] = '0'
		}
	}
	return string(st[:m])
}
