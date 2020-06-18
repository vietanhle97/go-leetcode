package Expression_Add_Operators_282

import "strconv"

func backtrack(res *[]string, num, path string, start, prev, sum, target int) {
	if sum == target && start >= len(num) {
		*res = append(*res, path)
		return
	}
	for i := start; i < len(num); i++ {
		cur := num[start : i+1]
		if len(cur) >= 2 && cur[0] == '0' {
			continue
		}
		if len(path) == 0 {
			val, _ := strconv.Atoi(cur)
			backtrack(res, num, path+cur, i+1, val, val, target)
		} else {
			val, _ := strconv.Atoi(cur)
			backtrack(res, num, path+"+"+cur, i+1, val, val+sum, target)
			backtrack(res, num, path+"*"+cur, i+1, prev*val, (sum-prev)+(prev*val), target)
			backtrack(res, num, path+"-"+cur, i+1, -val, sum-val, target)
		}
	}
}

func addOperators(num string, target int) []string {
	res := make([]string, 0)
	backtrack(&res, num, "", 0, 0, 0, target)
	return res
}
