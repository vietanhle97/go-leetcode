package Split_Two_Strings_To_Make_Palindrome_1616

func check(a, b string) int {
	n := len(a)
	i, j := 0, n-1
	for i < j {
		if a[i] != b[j] {
			return i
		}
		i++
		j--
	}
	return -1
}

func isPalin(a string) bool {
	i, j := 0, len(a)-1
	for i < j {
		if a[i] != a[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func checkPalindromeFormation(a string, b string) bool {
	n := len(a)
	pref, suff := check(a, b), check(b, a)
	if pref == -1 || suff == -1 {
		return true
	}
	s1, s2, s3, s4 := a[pref:n-pref], b[pref:n-pref], a[suff:n-suff], b[suff:n-suff]
	return isPalin(s1) || isPalin(s2) || isPalin(s3) || isPalin(s4)
}
