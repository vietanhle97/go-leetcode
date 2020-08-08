package Can_Convert_String_In_K_Moves_1540

func canConvertString(s string, t string, k int) bool {
	m, n := len(s), len(t)
	if m != n {
		return false
	}
	arr := make([]int, m)
	check := map[int]int{}
	for i := range s {
		tmp := int(t[i]) - int(s[i])
		if tmp == 0 {
			continue
		}
		if tmp < 0 {
			tmp += 26
		}
		if _, ok := check[tmp]; !ok {
			arr[i] = tmp
			check[tmp] = 1
		} else {
			cur := check[tmp]
			check[tmp]++
			tmp = 26*cur + tmp
			arr[i] = tmp

		}
	}
	// fmt.Println(arr)
	ans := arr[0]
	for _, e := range arr {
		if e > ans {
			ans = e
		}
	}
	return ans <= k
}
