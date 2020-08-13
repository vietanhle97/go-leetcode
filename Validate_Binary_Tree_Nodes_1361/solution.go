package Validate_Binary_Tree_Nodes_1361

func validateBinaryTreeNodes(n int, left []int, right []int) bool {
	if n == 1 {
		return true
	}
	parent := make([]int, n)
	cnt := 0
	root := -1
	for i := range parent {
		parent[i] = -1
	}
	for i, e := range left {
		if e != -1 {
			if parent[e] != -1 {
				return false
			}
			parent[e] = i
		}
	}
	for i, e := range right {
		if e != -1 {
			if parent[e] != -1 {
				return false
			}
			parent[e] = i
		}
	}
	for i, e := range parent {
		if e == -1 {
			cnt++
			root = i
		}
		if cnt > 1 {
			return false
		}
	}

	return cnt != 0 && (left[root] != -1 || right[root] != -1)
}
