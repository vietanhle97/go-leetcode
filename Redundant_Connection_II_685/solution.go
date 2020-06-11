package Redundant_Connection_II_685

func find(parent []int, i int) int {
	if parent[i] == -1 {
		return i
	}
	return find(parent, parent[i])
}

func union(parent []int, x, y int) {

}

func findRedundantDirectedConnection(edges [][]int) []int {

}
