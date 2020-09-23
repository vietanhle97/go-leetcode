package Make_Sum_Divisible_By_P_1590

const MaxInt = 1<<31 - 1

func bs(l, r, m int, arr []int) int {
	if l > r {
		return -1
	}
	mid := l + (r-l)/2
	rs := bs(mid+1, r, m, arr)
	if rs != -1 {
		return rs
	}
	if arr[mid]+1 <= m {
		return arr[mid] + 1
	}
	return bs(l, mid-1, m, arr)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minSubarray(nums []int, n int) int {
	m := len(nums)
	if m == 1 {
		if nums[0]%n == 0 {
			return 0
		}
		return -1
	}
	res := MaxInt
	p, s := []int{nums[0]}, nums[m-1]
	if p[0]%n == 0 || s%n == 0 {
		res = min(res, m-1)
	}
	S := map[int][]int{}
	S[s%n] = []int{0}
	for i := 1; i < m; i++ {
		p = append(p, p[i-1]+nums[i])
		s += nums[m-1-i]
		if p[i]%n == 0 || s%n == 0 {
			res = min(res, m-i-1)
		}
		if i == m-1 {
			if p[i]%n == 0 {
				return 0
			}
			if p[i] < n {
				return -1
			}
		}
		S[s%n] = append(S[s%n], i)
	}
	// find i,j that (p[i]+s[j])%p == n-r && i+j max
	for i, e := range p {
		if _, ok := S[n-e%n]; !ok {
			continue
		}
		j := bs(0, len(S[n-e%n])-1, m-1-i, S[n-e%n])
		if j != -1 {
			res = min(res, m-i-1-j)
		}
	}
	if res == MaxInt {
		return -1
	}
	return res
}
