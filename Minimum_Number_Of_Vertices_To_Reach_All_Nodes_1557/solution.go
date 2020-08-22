package Minimum_Number_Of_Vertices_To_Reach_All_Nodes_1557

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	indeg := make([]int, n)
	for _, e := range edges {
		indeg[e[1]]++
	}
	res := make([]int, 0)
	for i := range indeg {
		if indeg[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}
