package Split_Array_Into_Fibonacci_Sequence_842

import "strconv"

func backtrack(start int, S string, path []int, res *[]int) {
	if start >= len(S) && len(path) > 2 {
		*res = path
		return
	}
	for i := start; i < len(S); i++ {
		cur := S[start : i+1]

		if len(cur) >= 2 && cur[0] == '0' {
			break
		}

		tmp, e := strconv.ParseInt(cur, 10, 32)
		if e != nil {
			break
		}
		num := int(tmp)
		if len(path) < 2 {
			backtrack(i+1, S, append(path, num), res)
		} else {
			if int(tmp) == path[len(path)-1]+path[len(path)-2] {
				backtrack(i+1, S, append(path, num), res)
			}
		}
	}
}

func splitIntoFibonacci(S string) []int {
	res := make([]int, 0)
	backtrack(0, S, []int{}, &res)
	return res
}
