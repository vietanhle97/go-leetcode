package Word_Search_79

func isUsed(used [][]int, x, y int) bool {
	for i := range used {
		if used[i][0] == x && used[i][1] == y {
			return true
		}
	}
	return false
}

func DFS(x, y, m, n int, word string, used [][]int, board [][]byte) bool {
	if len(word) == 0 {
		return true
	}
	if isUsed(used, x, y) || x < 0 || x >= m || y < 0 || y >= n || board[x][y] != word[0] {
		return false
	}
	return DFS(x-1, y, m, n, word[1:], append(used, []int{x, y}), board) || DFS(x+1, y, m, n, word[1:], append(used, []int{x, y}), board) || DFS(x, y-1, m, n, word[1:], append(used, []int{x, y}), board) || DFS(x, y+1, m, n, word[1:], append(used, []int{x, y}), board)
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	if n == 0 {
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if DFS(i, j, m, n, word, [][]int{}, board) {
				return true
			}
		}
	}
	return false
}
