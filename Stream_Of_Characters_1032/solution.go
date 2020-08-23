package Stream_Of_Characters_1032

type Trie struct {
	child [26]*Trie
	isEnd bool
}

func convert(ch byte) int {
	return int(ch) - int('a')
}

func (root *Trie) insert(w string) {
	tmp, m := root, len(w)
	for i := m - 1; i >= 0; i-- {
		ind := convert(w[i])
		if tmp.child[ind] == nil {
			tmp.child[ind] = &Trie{}
		}
		tmp = tmp.child[ind]
	}
	tmp.isEnd = true
}

func (root *Trie) searchPrefix(w string) bool {
	tmp, m := root, len(w)
	for i := 0; i < m; i++ {
		ind := convert(w[i])
		if tmp.child[ind] == nil {
			return false
		}
		tmp = tmp.child[ind]
		if tmp.isEnd == true {
			return true
		}
	}
	return tmp != nil && tmp.isEnd
}

type StreamChecker struct {
	trie *Trie
	q    string
}

func Constructor(words []string) StreamChecker {
	root := &Trie{}
	for _, w := range words {
		root.insert(w)
	}
	return StreamChecker{root, ""}
}

func (check *StreamChecker) Query(letter byte) bool {
	check.q = string(letter) + check.q
	return check.trie.searchPrefix(check.q)
}
