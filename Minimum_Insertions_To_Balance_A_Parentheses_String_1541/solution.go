package Minimum_Insertions_To_Balance_A_Parentheses_String_1541

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func minInsertions(s string) int {
	ans, bal := 0, 0
	for i := range s {
		if s[i] == '(' {
			if bal > 0 && bal%2 == 1 {
				ans += 1
				bal -= 1
			} else if bal < 0 && bal%2 == -1 {
				ans += abs(bal-1)/2 + 1
				bal = 0
			}
			bal += 2
		} else {
			bal -= 1
		}
		// fmt.Println(bal)
		if bal == -2 {
			ans += 1
			bal += 2
		}
	}
	if bal == -1 {
		return ans + 2
	}
	return ans + bal
}
