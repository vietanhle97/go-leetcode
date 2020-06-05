package Vertical_Order_Traversal_Of_Binary_Tree_987

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Level struct {
	m map[int]map[int][]int
}

func findOrder(node *TreeNode, lvl *Level, x int, y int) {
	if node == nil {
		return
	}
	findOrder(node.Left, lvl, x-1, y+1)
	findOrder(node.Right, lvl, x+1, y+1)
	if _, ok := lvl.m[x]; !ok {
		lvl.m[x] = map[int][]int{}
		lvl.m[x][y] = []int{node.Val}
	} else {
		lvl.m[x][y] = append(lvl.m[x][y], node.Val)
	}
	sort.Ints(lvl.m[x][y])
}

func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	lvl := Level{map[int]map[int][]int{}}
	findOrder(root, &lvl, 0, 0)
	res := make([][]int, 0)
	var keys []int
	for k := range lvl.m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, i := range keys {
		var tmp []int
		for k := range lvl.m[i] {
			tmp = append(tmp, k)
		}
		sort.Ints(tmp)
		for _, j := range tmp {
			cur := []int{}
			for _, e := range lvl.m[i][j] {
				cur = append(cur, e)
			}
			res = append(res, cur)
		}
	}
	return res
}
