package Ways_To_Split_Array_Into_Three_Subarrays_1712

const mod = int(1e9 + 7)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func bsLeft(pref []int, cur, l, r int) int {
	res := -1
	for l <= r {
		mid := l + (r-l)/2
		tmp := pref[mid] - cur
		if tmp >= cur {
			if pref[len(pref)-1]-pref[mid] >= tmp {
				if res == -1 {
					res = mid
				}
				res = min(res, mid)
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}

func bsRight(pref []int, cur, l, r int) int {
	res := -1
	for l <= r {
		mid := l + (r-l)/2
		tmp := pref[mid] - cur
		if tmp >= cur {
			if pref[len(pref)-1]-pref[mid] >= tmp {
				if res == -1 {
					res = mid
				}
				res = max(res, mid)
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			l = mid + 1
		}
	}
	return res
}

func waysToSplit(nums []int) int {
	pref := []int{0}
	sum := 0
	m := len(nums)
	cnt, res := 1, 0
	for _, e := range nums {
		pref = append(pref, pref[len(pref)-1]+e)
		sum = pref[len(pref)-1]
	}

	for cnt < m && pref[cnt] <= sum/3 {
		l, r := bsLeft(pref, pref[cnt], cnt+1, m), bsRight(pref, pref[cnt], cnt+1, m)
		if r >= l && r != -1 && l != -1 {
			if sum != 0 {
				res += r - l + 1
			} else {
				res += r - l
			}
			res %= mod
		}
		cnt++
	}
	return res
}
