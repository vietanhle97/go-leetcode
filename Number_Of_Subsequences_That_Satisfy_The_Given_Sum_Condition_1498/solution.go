package Number_Of_Subsequences_That_Satisfy_The_Given_Sum_Condition_1498

import "sort"

func pow2(a int) []int {
	MOD := int(1e9 + 7)
	cnt := 0
	prod := 1
	res := make([]int, a)
	for cnt < a {
		res[cnt] = prod
		prod = 2 * prod % MOD
		cnt++
	}
	return res
}

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	m := len(nums)
	res := make([]int, m)
	for i := 0; i < m; i++ {
		if nums[i]+nums[0] <= target {
			res[0] = i
		}
	}

	for i := 1; i < m; i++ {
		if nums[i]*2 > target {
			break
		}
		if nums[res[i-1]]+nums[i] <= target {
			res[i] = res[i-1]
		} else {
			j := res[i-1] - 1
			for j >= 0 {
				if nums[j]+nums[i] <= target {
					res[i] = j
					break
				}
				j--
			}
		}
	}

	ans := 0
	pow := pow2(res[0] + 1)
	MOD := int(1e9 + 7)
	for i, e := range res {
		if 2*nums[i] <= target {
			ans += pow[e-i]
			ans %= MOD
		}
	}
	return ans % MOD
}
