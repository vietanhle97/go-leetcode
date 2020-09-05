package Number_Of_Ways_To_Split_A_String_1573

const mod = int(1e9) + 7

func numWays(s string) int {
	n := len(s)
	num := make([]int, len(s))
	last := 0
	for i := range s {
		if s[i] == '1' {
			last++
			num[i] = last
		}
	}
	if last%3 != 0 {
		return 0
	}
	if last == 0 {
		return ((n - 1) * (n - 2) / 2) % mod
	}
	target := last / 3
	firstEnd, midStart, midEnd, lastStart := 0, 0, 0, 0
	for i, e := range num {
		if e == target {
			firstEnd = i
		}
		if e == target+1 {
			midStart = i
		}
		if e == 2*target {
			midEnd = i
		}
		if e == 2*target+1 {
			lastStart = i
		}
	}

	return (((midStart - firstEnd) % mod) * ((lastStart - midEnd) % mod)) % mod
}
