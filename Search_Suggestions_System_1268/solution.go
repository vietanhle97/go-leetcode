package Search_Suggestions_System_1268

import "sort"

func binarySearch(products []string, s string, l, r int) int {
	if l > r {
		return -1
	}
	n, mid := len(s), l+(r-l)/2
	cur := products[mid]
	ls := binarySearch(products, s, l, mid-1)
	if ls != -1 {
		return ls
	}
	if len(cur) >= len(s) && cur[0:n] == s {
		return mid
	}
	return binarySearch(products, s, mid+1, r)
}

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	res := make([][]string, 0)
	n := len(products)
	for i := range searchWord {
		tmp := make([]string, 0)
		s := searchWord[0 : i+1]
		l, r := 0, n-1
		for len(tmp) < 3 {
			ind := binarySearch(products, s, l, r)
			if ind == -1 {
				break
			}
			// fmt.Println(ind,products[ind])
			l = ind + 1
			tmp = append(tmp, products[ind])
		}
		res = append(res, tmp)
	}
	return res
}
