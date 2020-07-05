package Kth_Ancestor_Of_A_Tree_Node_1483

type TreeAncestor struct {
	n  int
	sp [][]int
}

func Constructor(n int, parent []int) TreeAncestor {
	sp := make([][]int, n)
	for i := 0; i < n; i++ {
		sp[i] = make([]int, 17)
		sp[i][0] = parent[i]
	}

	for j := 1; j < 17; j++ {
		for i := 0; i < n; i++ {
			p := sp[i][j-1]
			if p != -1 {
				sp[i][j] = sp[p][j-1]
			} else {
				sp[i][j] = -1
			}
		}
	}
	return TreeAncestor{n, sp}
}

func (anc *TreeAncestor) GetKthAncestor(node int, k int) int {
	for i := 0; i < 17; i++ {
		if k&1 == 1 {
			node = anc.sp[node][i]
			if node == -1 {
				return -1
			}
		}
		k = k >> 1
	}
	return node
}
