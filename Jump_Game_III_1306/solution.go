package Jump_Game_III_1306

func dfs(v int, arr []int, vis *[]bool, res *bool) {
	if arr[v] == 0 || *res {
		*res = true
		return
	}
	(*vis)[v] = true
	if v-arr[v] >= 0 && (*vis)[v-arr[v]] == false {
		dfs(v-arr[v], arr, vis, res)
	}
	if v+arr[v] < len(arr) && (*vis)[v+arr[v]] == false {
		dfs(v+arr[v], arr, vis, res)
	}
}
func canReach(arr []int, start int) bool {
	vis := make([]bool, len(arr))
	res := false
	dfs(start, arr, &vis, &res)
	return res
}
