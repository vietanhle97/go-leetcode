package Group_Anagrams_49

func isAnagram(s1, s2 string) bool {
	if len(s1) == len(s2) && len(s1) == 0 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	arr := make([]int, 26)
	for i := range s1 {
		arr[int(s1[i]-'a')]++
	}
	for i := range s2 {
		arr[int(s2[i]-'a')]--
		if arr[int(s2[i]-'a')] < 0 {
			return false
		}
	}
	return true
}
func groupAnagrams(strs []string) [][]string {

	res := make([][]string, 0)
	n := len(strs)
	vis := map[int]bool{}
	for i := 0; i < n; i++ {
		tmp := make([]string, 0)
		if _, ok := vis[i]; !ok {
			tmp = append(tmp, strs[i])
		}
		for j := i; j < n; j++ {
			if i != j {
				if _, ok := vis[j]; !ok {
					if isAnagram(strs[i], strs[j]) {
						tmp = append(tmp, strs[j])
						vis[j] = true
					}
				}
			}
		}
		if len(tmp) > 0 {
			res = append(res, tmp)
		}
	}
	return res
}
