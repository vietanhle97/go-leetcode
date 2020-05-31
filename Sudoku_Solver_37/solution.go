package Sudoku_Solver_37

func findEmpty(board [][]byte) (bool, int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == byte('.') {
				return true, i, j
			}
		}
	}
	return false, -1, -1
}

func isValidInRow(row int, n byte, board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == n {
			return false
		}
	}
	return true
}

func isValidInCol(col int, n byte, board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == n {
			return false
		}
	}
	return true
}

func isValidInBox(row int, col int, n byte, board [][]byte) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[row+i][col+j] == n {
				return false
			}
		}
	}
	return true
}

func isValid(row int, col int, n byte, board [][]byte) bool {
	return isValidInRow(row, n, board) && isValidInCol(col, n, board) && isValidInBox(row-row%3, col-col%3, n, board)
}

func backtrack(board [][]byte) bool {
	ok, _, _ := findEmpty(board)
	if !ok {
		return true
	}
	for i := 1; i < 10; i++ {
		_, row, col := findEmpty(board)
		cur := []byte(string(i + 48))[0]
		if isValid(row, col, cur, board) {
			board[row][col] = cur
			if backtrack(board) {
				return true
			}
		}
		board[row][col] = byte('.')
	}
	return false
}

func solveSudoku(board [][]byte) {
	backtrack(board)
}
