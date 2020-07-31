package Implement_Magic_Dictionary_676

type MagicDictionary struct {
	children [26]*MagicDictionary
	isEnd    bool
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{}
}

/** Build a dictionary through a list of words */
func (magic *MagicDictionary) BuildDict(dict []string) {
	cur := magic
	for _, word := range dict {
		l := len(word)
		tmp := cur
		for i := 0; i < l; i++ {
			ind := int(word[i] - 'a')
			if tmp.children[ind] == nil {
				tmp.children[ind] = &MagicDictionary{}
			}
			tmp = tmp.children[ind]
		}
		tmp.isEnd = true
	}

}

func (magic *MagicDictionary) search(cnt int, word string) (int, bool, string) {
	l := len(word)
	cur := magic
	for i := 0; i < l; i++ {
		ind := int(word[i] - 'a')
		tmp, end, w := 0, false, ""
		for j := 0; j < 26; j++ {
			if cur.children[j] != nil {
				if j == ind {
					tmp, end, w = cur.children[j].search(cnt, word[i+1:])
					if tmp == 1 && end && len(w) == 0 {
						return 1, true, w
					}
				} else {
					tmp, end, w = cur.children[j].search(cnt+1, word[i+1:])
					if tmp == 1 && end && len(w) == 0 {
						return 1, true, w
					}
				}
			}
		}
		if !end || tmp > 1 || len(w) != 0 {
			return cnt, false, word
		}
	}
	if cnt == 1 {
		return 1, cur != nil && cur.isEnd, word
	}
	return cnt, cur != nil && cur.isEnd, word
}

/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (magic *MagicDictionary) Search(word string) bool {
	cnt, b, _ := magic.search(0, word)
	return cnt == 1 && b
}
