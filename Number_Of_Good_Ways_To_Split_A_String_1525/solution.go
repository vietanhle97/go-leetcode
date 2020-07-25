package Number_Of_Good_Ways_To_Split_A_String_1525

func numSplits(s string) int {
	n := len(s)
	diff1 := make([]int, n)
	diff2 := make([]int, n)
	m1, m2 := [26]int{}, [26]int{}
	for i := range s {
		if i == 0 {
			diff1[i] = 1
			m1[s[i]-'a']++
		} else {
			if m1[s[i]-'a'] == 0 {
				diff1[i] = diff1[i-1] + 1
				m1[s[i]-'a']++
			} else {
				diff1[i] = diff1[i-1]

			}
		}
	}
	cnt := 0
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			diff2[i] = 1
			m2[s[i]-'a']++
		} else {
			if m2[s[i]-'a'] == 0 {
				diff2[i] = diff2[i+1] + 1
				m2[s[i]-'a']++
			} else {
				diff2[i] = diff2[i+1]

			}
		}
	}

	for i := 0; i < n-1; i++ {
		if diff1[i] == diff2[i+1] {
			cnt++
		}
	}
	// fmt.Println(diff1, diff2)
	return cnt
}
