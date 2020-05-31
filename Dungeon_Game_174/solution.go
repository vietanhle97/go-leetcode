package Dungeon_Game_174

import (
	"strconv"
)

type Mem struct {
	m map[string]int
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func dfs(row int, col int, grid [][]int, mem *Mem) int {
	tmp := "(" + strconv.Itoa(row) + "," + strconv.Itoa(col) + ")"
	if _, ok := mem.m[tmp]; ok {
		return mem.m[tmp]
	}
	if row == len(grid)-1 && col == len(grid[0])-1 {
		return grid[row][col]
	}
	k := -99999999999
	step := [][]int{{row + 1, col}, {row, col + 1}}

	for _, y := range step {
		if y[0] >= len(grid) || y[1] >= len(grid[0]) {
			continue
		}
		k = max(k, dfs(y[0], y[1], grid, mem))
	}
	t := grid[row][col]
	if k < 0 {
		t += k
	}
	if t > 0 {
		t = 0
	}
	mem.m[tmp] = t
	return mem.m[tmp]
}

func calculateMinimumHP(dungeon [][]int) int {
	mem := Mem{map[string]int{}}
	res := dfs(0, 0, dungeon, &mem) - 1
	if res < 0 {
		return -res
	}
	return 1
}
