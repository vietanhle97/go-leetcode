package Number_Of_Sub_arrays_With_Odd_Sum_1524

func numOfSubarrays(arr []int) int {
	MOD := int(1e9) + 7
	tmp := []int{1, 0}
	res, val := 0, 0
	for i := range arr {
		val = ((val+arr[i])%2 + 2) % 2
		tmp[val]++
	}
	res = tmp[0] * tmp[1]
	return res % MOD
}
