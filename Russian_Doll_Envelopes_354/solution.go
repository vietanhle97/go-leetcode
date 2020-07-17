package Russian_Doll_Envelopes_354

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxEnvelopes(env [][]int) int {
	n := len(env)
	if n == 0 {
		return 0
	}
	sort.Slice(env, func(i, j int) bool {
		if env[i][0] == env[j][0] {
			return env[i][1] > env[j][1]
		} else {
			return env[i][0] < env[j][0]
		}
	})
	cnt := make([]int, n)
	cnt[0] = 1
	res := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if env[i][0] > env[j][0] && env[i][1] > env[j][1] {
				cnt[i] = max(cnt[i], cnt[j]+1)
				res = max(res, cnt[i])
			} else {
				cnt[i] = max(cnt[i], 1)
				res = max(res, cnt[i])
			}
		}
	}
	// fmt.Println(env, cnt)
	return res
}
