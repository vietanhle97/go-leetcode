package Implement_TrieNode_Prefix_Tree_208

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

func charToIndex(ch byte) int {
	return int(ch - 'a')
}

func (root *TrieNode) searchPrefix(s string) *TrieNode {
	tmp := root
	l := len(s)
	for i := 0; i < l; i++ {
		ind := charToIndex(s[i])
		if tmp.children[ind] == nil {
			return nil
		}
		tmp = tmp.children[ind]
	}
	return tmp
}

/** Initialize your data structure here. */
func Constructor() TrieNode {
	return TrieNode{}
}

/** Inserts a word into the trie. */
func (root *TrieNode) Insert(word string) {
	tmp := root
	l := len(word)
	for i := 0; i < l; i++ {
		ind := charToIndex(word[i])
		if tmp.children[ind] == nil {
			tmp.children[ind] = &TrieNode{}
		}
		tmp = tmp.children[ind]
	}
	tmp.isEnd = true
}

/** Returns if the word is in the trie. */
func (root *TrieNode) Search(word string) bool {
	tmp := root
	l := len(word)
	for i := 0; i < l; i++ {
		ind := charToIndex(word[i])
		if tmp.children[ind] == nil {
			return false
		}
		tmp = tmp.children[ind]
	}
	return tmp != nil && tmp.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (root *TrieNode) StartsWith(prefix string) bool {
	tmp := root.searchPrefix(prefix)
	return tmp != nil
}
