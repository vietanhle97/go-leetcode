package Prefix_And_Suffix_Search_745

type TrieNode struct {
	children [27]*TrieNode
	w        int
}

type WordFilter struct {
	trie *TrieNode
}

func Constructor(words []string) WordFilter {
	trie := &TrieNode{}
	for w := range words {
		word := words[w] + "{"
		l := len(word)
		for i := 0; i < l; i++ {
			cur := trie
			cur.w = w
			for j := i; j < 2*l-1; j++ {
				k := word[j%l] - 'a'
				if cur.children[k] == nil {
					cur.children[k] = &TrieNode{}
				}
				cur = cur.children[k]
				cur.w = w
			}
		}
	}
	return WordFilter{trie}
}

func (filter *WordFilter) F(prefix string, suffix string) int {
	cur := filter.trie
	tmp := suffix + "{" + prefix
	for _, e := range tmp {
		if cur.children[e-'a'] == nil {
			return -1
		}
		cur = cur.children[e-'a']

	}
	return cur.w
}
