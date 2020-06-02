package Count_Numbers_with_Unique_Digits_357

func countNumbersWithUniqueDigits(n int) int {
	res := 0
	if n == 0 {
		return 1
	}
	for n > 0 {
		cnt := 1
		for i := 0; i < n; i++ {
			if n == 1 {
				cnt *= 10
			} else if i == 0 {
				cnt *= 9
			} else {
				cnt *= 10 - i
			}
		}
		res += cnt
		n--
	}
	return res
}
