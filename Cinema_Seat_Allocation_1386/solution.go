package Cinema_Seat_Allocation_1386

import "sort"

func maxNumberOfFamilies(n int, rs [][]int) int {
	m := len(rs)
	if m == 0 {
		return n * 2
	}

	sort.Slice(rs, func(i, j int) bool {
		if rs[i][0] == rs[j][0] {
			return rs[i][1] < rs[j][1]
		}
		return rs[i][0] < rs[j][0]
	})
	avai := []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 2}
	cnt := 0
	res := map[int][]int{}
	cur := rs[0][0]
	for i, e := range rs {
		if e[0] > cur {
			cnt += avai[10-res[cur][len(res[cur])-1]]
			cur = e[0]
		}
		if _, ok := res[e[0]]; !ok {
			if e[1] > 5 {
				cnt += avai[e[1]-1]
			}
			res[e[0]] = append(res[e[0]], e[1])
		} else {
			tmp := e[1] - res[e[0]][len(res[e[0]])-1]
			if tmp > 5 {
				cnt += avai[tmp]
			} else if tmp == 5 {
				if e[1] != 7 && e[1] != 9 {
					cnt += avai[tmp]
				}
			}

			res[e[0]] = append(res[e[0]], e[1])
		}
		if i == m-1 {
			cnt += avai[10-res[e[0]][len(res[e[0]])-1]]
		}
	}
	return cnt + 2*(n-len(res))
}
