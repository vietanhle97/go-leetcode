package Word_Search_II_212

type Trie struct {
	children [26]*Trie
	word     string
}

func charToIndex(ch byte) int {
	return int(ch - 'a')
}

/** Inserts a list of words into the trie. */
func Insert(words []string) *Trie {
	root := &Trie{}
	for _, word := range words {
		tmp := root
		l := len(word)
		for i := 0; i < l; i++ {
			ind := charToIndex(word[i])
			if tmp.children[ind] == nil {
				tmp.children[ind] = &Trie{}
			}
			tmp = tmp.children[ind]
		}
		tmp.word = word
	}
	return root
}

func DFS(x, y int, board [][]byte, root *Trie, res *[]string) {
	ch := board[x][y]
	if ch == '.' || (*root).children[ch-'a'] == nil {
		return
	}
	root = root.children[ch-'a']
	if len(root.word) != 0 {
		*res = append(*res, root.word)
		root.word = ""
	}
	board[x][y] = '.'
	if x > 0 {
		DFS(x-1, y, board, root, res)
	}
	if y > 0 {
		DFS(x, y-1, board, root, res)
	}
	if x < (len(board) - 1) {
		DFS(x+1, y, board, root, res)
	}
	if y < len(board[0])-1 {
		DFS(x, y+1, board, root, res)
	}
	board[x][y] = ch
}

func findWords(board [][]byte, words []string) []string {
	root := Insert(words)
	m, n, res := len(board), len(board[0]), make([]string, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			DFS(i, j, board, root, &res)
		}
	}
	return res
}
