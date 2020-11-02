package Repeated_DNA_Sequences_187

func findRepeatedDnaSequences(s string) []string {
	m := map[string]int{}
	init, res := "", []string{}
	if len(s) < 10 {
		return res
	}
	init = s[:10]
	m[init] = 1
	for i := 10; i < len(s); i++ {
		init = init[1:] + string(s[i])
		m[init]++
		if m[init] == 2 {
			res = append(res, init)
		}
	}
	return res
}
