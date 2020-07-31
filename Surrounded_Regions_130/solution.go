package Surrounded_Regions_130

func dfs(m, n, i, j int, g [][]byte, vis [][]bool) {
	if vis[i][j] {
		return
	}
	vis[i][j] = true
	if i > 0 && !vis[i-1][j] && g[i-1][j] == 'O' {
		dfs(m, n, i-1, j, g, vis)
	}
	if j > 0 && !vis[i][j-1] && g[i][j-1] == 'O' {
		dfs(m, n, i, j-1, g, vis)
	}
	if i < m-1 && !vis[i+1][j] && g[i+1][j] == 'O' {
		dfs(m, n, i+1, j, g, vis)
	}
	if j < n-1 && !vis[i][j+1] && g[i][j+1] == 'O' {
		dfs(m, n, i, j+1, g, vis)
	}
}

func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	if n == 0 {
		return
	}
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			dfs(m, n, i, 0, board, vis)
		}
		if board[i][n-1] == 'O' {
			dfs(m, n, i, n-1, board, vis)
		}
	}
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			dfs(m, n, 0, j, board, vis)
		}
		if board[m-1][j] == 'O' {
			dfs(m, n, m-1, j, board, vis)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !vis[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}
