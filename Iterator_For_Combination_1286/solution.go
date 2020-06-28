package Iterator_For_Combination_1286

type CombinationIterator struct {
	data []string
}

func backtrack(start, length int, s, path string, all *[]string) {
	if len(path) == length {
		*all = append(*all, path)
		return
	}
	for i := start; i < len(s); i++ {
		if len(path) < length {
			backtrack(i+1, length, s, path+string(s[i]), all)
		}
	}
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	com := CombinationIterator{[]string{}}
	backtrack(0, combinationLength, characters, "", &com.data)
	return com
}

func (this *CombinationIterator) Next() string {

	res := this.data[0]
	this.data = this.data[1:]
	return res
}

func (this *CombinationIterator) HasNext() bool {
	return len(this.data) > 0
}
