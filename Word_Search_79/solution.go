package Word_Search_79

func isUsed(used [][]bool, x, y int) bool {
	return used[x][y]
}

func DFS(x, y, m, n int, word string, used [][]bool, board [][]byte) bool {
	if len(word) == 0 {
		return true
	}
	if x < 0 || x >= m || y < 0 || y >= n || isUsed(used, x, y) || board[x][y] != word[0] {
		return false
	}
	used[x][y] = true
	if DFS(x-1, y, m, n, word[1:], used, board) {
		return true
	}

	if DFS(x+1, y, m, n, word[1:], used, board) {
		return true
	}

	if DFS(x, y-1, m, n, word[1:], used, board) {
		return true
	}

	if DFS(x, y+1, m, n, word[1:], used, board) {
		return true
	}

	used[x][y] = false
	return false
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
	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if DFS(i, j, m, n, word, used, board) {
				return true
			}
		}
	}
	return false
}
