package Palindrome_Pairs_336

type Trie struct {
	child []*Trie
	isEnd bool
	ind   int
	palin []int
}

func isPalindrome(l, r int, s string) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func charToInd(ch byte) int {
	return int(ch) - int('a')
}

func initTrie() *Trie {
	child := make([]*Trie, 26)
	ind := -1
	palin := make([]int, 0)
	return &Trie{child, false, ind, palin}
}

func (root *Trie) Insert(w string, k int) {
	tmp, m := root, len(w)
	for i := m - 1; i >= 0; i-- {
		ind := charToInd(w[i])
		if tmp.child[ind] == nil {
			tmp.child[ind] = initTrie()
		}
		if isPalindrome(0, i, w) {
			tmp.palin = append(tmp.palin, k)
		}
		tmp = tmp.child[ind]
	}
	tmp.isEnd = true
	tmp.palin = append(tmp.palin, k)
	tmp.ind = k
}

func (root *Trie) Search(w string, cur int, res *[][]int) {
	tmp, m := root, len(w)
	for i := 0; i < m; i++ {
		ind := charToInd(w[i])
		if tmp.ind >= 0 && tmp.ind != cur && isPalindrome(i, m-1, w) {
			*res = append(*res, []int{cur, tmp.ind})
		}
		if tmp.child[ind] == nil {
			return
		}
		tmp = tmp.child[ind]
	}
	for _, e := range tmp.palin {
		if e != cur {
			*res = append(*res, []int{cur, e})
		}
	}
}

func palindromePairs(words []string) [][]int {
	root, res := initTrie(), make([][]int, 0)
	for i, w := range words {
		root.Insert(w, i)
	}
	for i, w := range words {
		root.Search(w, i, &res)
	}
	return res
}
