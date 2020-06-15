package Unique_Binary_Search_Trees_96

func numTrees(n int) int {
	table := make([]int, n+1)
	table[0] = 1
	table[1] = 1
	for i := 2; i < n+1; i++ {
		for j := 1; j < i+1; j++ {
			table[i] += table[i-j] * table[j-1]
		}
	}
	return table[n]
}
