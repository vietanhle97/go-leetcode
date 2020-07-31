package Add_And_Search_Word_Data_Structure_Design_211

type WordDictionary struct {
	children [27]*WordDictionary
	isEnd    bool
}

func charToIndex(ch byte) int {
	if ch == '.' {
		return 26
	}
	return int(ch - 'a')
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

/** Adds a word into the data structure. */
func (root *WordDictionary) AddWord(word string) {
	tmp := root
	l := len(word)
	for i := 0; i < l; i++ {
		ind := charToIndex(word[i])
		if tmp.children[ind] == nil {
			tmp.children[ind] = &WordDictionary{}
		}
		tmp = tmp.children[ind]
	}
	tmp.isEnd = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (root *WordDictionary) Search(word string) bool {
	tmp := root
	l := len(word)
	for i := 0; i < l; i++ {
		ind := charToIndex(word[i])
		if ind == 26 {
			for j := 0; j < 27; j++ {
				if tmp.children[j] != nil {
					res := tmp.children[j].Search(word[i+1:])
					if res {
						return true
					}
				}
			}
			return false
		} else {
			if tmp.children[ind] == nil {
				return false
			}
			tmp = tmp.children[ind]
		}
	}
	return tmp != nil && tmp.isEnd
}
