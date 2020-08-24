package Minimum_Add_To_Make_Parentheses_Valid_921

func minAddToMakeValid(s string) int {
	bal, cnt := 0, 0
	for _, e := range s {
		if e == '(' {
			bal++
		} else {
			bal--
		}
		if bal == -1 {
			cnt++
			bal++
		}
	}
	return cnt + bal
}
