package Maximum_Sum_Of_3_Non_Overlapping_Subarrays_689

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n, init := len(nums), 0
	for i := 0; i < k; i++ {
		init += nums[i]
	}
	pref := []int{init}
	for i := k; i < n; i++ {
		init = init - nums[i-k] + nums[i]
		pref = append(pref, init)
	}
	res := []int{0, 0, 0}
	m := len(pref)
	cur := -1
	ls, rs := make([]int, m), make([]int, m)
	l, r := 0, m-1
	for i := 0; i < m; i++ {
		if pref[i] > pref[l] {
			l = i
		}
		ls[i] = l
	}
	for i := m - 1; i >= 0; i-- {
		if pref[i] >= pref[r] {
			r = i
		}
		rs[i] = r
	}
	for b := k; b < m-k; b++ {
		a, c := ls[b-k], rs[b+k]
		sum := pref[a] + pref[b] + pref[c]
		if cur == -1 || sum > cur {
			cur = sum
			res = []int{a, b, c}
		}
	}
	return res
}
