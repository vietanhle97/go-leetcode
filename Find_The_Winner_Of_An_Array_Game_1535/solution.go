package Find_The_Winner_Of_An_Array_Game_1535

func findMax(arr []int) int {
	res := arr[0]
	for _, e := range arr {
		if e > res {
			res = e
		}
	}
	return res
}
func getWinner(arr []int, k int) int {
	MAX := findMax(arr)
	m := len(arr)
	if arr[0] == MAX {
		return arr[0]
	} else if k >= m {
		return MAX
	}
	cnt, st := 0, []int{arr[0]}
	arr = arr[1:]
	for {
		if st[0] > arr[0] {
			cnt++
			tmp := arr[0]
			arr = arr[1:]
			arr = append(arr, tmp)
		} else {
			st[0] = arr[0]
			tmp := arr[0]
			arr = arr[1:]
			arr = append(arr, tmp)
			cnt = 1
		}
		if cnt == k {
			break
		}
	}

	return st[0]
}
