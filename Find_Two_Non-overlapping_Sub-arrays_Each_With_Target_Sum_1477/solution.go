package Find_Two_Non_overlapping_Sub_arrays_Each_With_Target_Sum_1477

const MaxInt = int(1e9)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func minSumOfLengths(arr []int, target int) int {
	if len(arr) == 1 {
		return -1
	}
	d := map[int]int{}
	d[0] = -1
	pref := []int{0}
	for i := range arr {
		pref = append(pref, pref[len(pref)-1]+arr[i])
		d[pref[len(pref)-1]] = i
	}
	res := MaxInt
	m, n := MaxInt, MaxInt
	for i := range arr {
		diff := pref[i+1] - target
		if _, ok := d[diff]; ok {
			m = min(m, i-d[diff])
		}
		if _, ok := d[pref[i+1]+target]; ok && m < MaxInt {
			n = d[pref[i+1]+target]
			res = min(res, m+n-i)
		}
	}
	if res >= 1e9 {
		return -1
	}
	return res
}
