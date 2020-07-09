package Reconstruct_Itinerary_332

import "sort"

func dfs(start string, m map[string][]string, ans *[]string) {
	for len(m[start]) > 0 {
		cur := m[start][0]
		m[start] = m[start][1:]
		dfs(cur, m, ans)
	}

	*ans = append(*ans, start)
}

func findItinerary(tickets [][]string) []string {
	m := make(map[string][]string, len(tickets)+1)
	var ans []string

	for _, t := range tickets {
		m[t[0]] = append(m[t[0]], t[1])
	}
	for k := range m {
		sort.Strings(m[k])
	}

	dfs("JFK", m, &ans)

	i, j := 0, len(ans)-1
	for i < j {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	return ans
}
