package All_Nodes_Distance_K_In_Binary_Tree_863

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func postOrder(root *TreeNode, dist *[]int) {
	if root == nil {
		return
	}
	postOrder(root.Left, dist)
	postOrder(root.Right, dist)
	if root.Left != nil {
		if (*dist)[root.Left.Val] != -1 {
			(*dist)[root.Val] = (*dist)[root.Left.Val] + 1
		}
	}
	if root.Right != nil {
		if (*dist)[root.Right.Val] != -1 {
			(*dist)[root.Val] = (*dist)[root.Right.Val] + 1
		}
	}
}

func preOrder(root *TreeNode, dist *[]int, d int) {
	if root == nil {
		return
	}
	if (*dist)[root.Val] == -1 {
		(*dist)[root.Val] = d
	}
	preOrder(root.Left, dist, min((*dist)[root.Val], d)+1)
	preOrder(root.Right, dist, min((*dist)[root.Val], d)+1)
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	dist := make([]int, 501)
	res := make([]int, 0)
	for i := range dist {
		dist[i] = -1
	}
	dist[target.Val] = 0
	postOrder(root, &dist)
	preOrder(root, &dist, dist[root.Val])
	for i := range dist {
		if dist[i] == K {
			res = append(res, i)
		}
	}
	return res
}
