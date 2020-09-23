package Robot_Bounded_In_Circle_1041

func isRobotBounded(ins string) bool {
	total := ins + ins + ins + ins
	cur := []int{0, 0}
	north, east, west, south := []int{0, -1}, []int{1, 0}, []int{-1, 0}, []int{0, 1}
	dirs := [][]int{north, east, south, west}
	dir := 0
	for _, e := range total {
		if e == 'G' {
			cur[0] += dirs[dir][0]
			cur[1] += dirs[dir][1]
		} else if e == 'R' {
			dir++
			dir %= 4
		} else {
			dir--
			if dir < 0 {
				dir += 4
			}
			dir %= 4
		}
	}
	if cur[0] == 0 && cur[1] == 0 && dir == 0 {
		return true
	}
	return false
}
