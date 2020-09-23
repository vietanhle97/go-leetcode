package Check_If_String_Is_Transformable_With_Substring_Sort_Operations_1585

import "strconv"

func isTransformable(s string, t string) bool {
	p := map[int][]int{}
	m := len(s)
	for i := m - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(string(s[i]))
		p[n] = append(p[n], i)
	}
	for _, e := range t {
		k, _ := strconv.Atoi(string(e))
		if _, ok := p[k]; !ok {
			return false
		}
		i := p[k][len(p[k])-1]
		for j := 0; j < k; j++ {
			if len(p[j]) > 0 && p[j][len(p[j])-1] < i {
				return false
			}
		}
		p[k] = p[k][:len(p[k])-1]
	}
	return true
}
