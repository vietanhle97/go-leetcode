package Gas_Station_134

func try(fuel, cur, m int, cost, gas []int, vis []bool, ind int) bool {
	// fmt.Println(cur, fuel, ind)
	if vis[cur] == true && fuel >= cost[cur] {
		return true
	}
	vis[cur] = true
	if cost[cur%m] > fuel {
		return false
	}
	return try(fuel+gas[(cur+1)%m]-cost[cur], (cur+1)%m, m, cost, gas, vis, ind)
}

func canCompleteCircuit(gas []int, cost []int) int {
	m := len(gas)
	for i := range gas {
		vis := make([]bool, m)
		if try(gas[i], i, m, cost, gas, vis, i) {
			return i
		}
	}
	return -1
}
