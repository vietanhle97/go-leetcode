package Find_The_Minimum_Number_Of_Fibonacci_Numbers_Whose_Sum_Is_K_1414

func findMinFibonacciNumbers(k int) int {
	fib := []int{1, 1}
	for {
		tmp := fib[len(fib)-1] + fib[len(fib)-2]
		if tmp > k {
			break
		}
		fib = append(fib, tmp)
	}
	j := len(fib) - 1
	cnt := 0
	for k != 0 {
		if k >= fib[j] {
			k -= fib[j]
			cnt++
		}
		j--
	}
	return cnt
}
