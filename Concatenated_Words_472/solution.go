package Concatenated_Words_472

import "sort"

type Trie struct {
	child [26]*Trie
	isEnd bool
	w     string
}

func convert(ch byte) int {
	return int(ch) - int('a')
}

func (root *Trie) insert(w string) {
	tmp, m := root, len(w)
	for i := 0; i < m; i++ {
		ind := convert(w[i])
		if tmp.child[ind] == nil {
			tmp.child[ind] = &Trie{}
		}
		tmp = tmp.child[ind]
	}
	tmp.w = w
	tmp.isEnd = true
}

func (root *Trie) search(w, p string) bool {
	tmp, m := root, len(w)
	if m == 0 {
		return true
	}
	for i := 0; i < m; i++ {
		ind := convert(w[i])
		if tmp.child[ind] == nil {
			return false
		}
		tmp = tmp.child[ind]
		if tmp.isEnd == true && tmp.w != p {
			if root.search(w[i+1:], w) {
				return true
			}
		}
	}
	return false
}

func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	root, res := &Trie{}, make([]string, 0)
	for _, w := range words {
		root.insert(w)
	}
	for _, w := range words {
		if len(w) != 0 && root.search(w, w) {
			res = append(res, w)
		}
	}
	return res
}
