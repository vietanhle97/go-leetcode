package Number_Of_Good_Leaf_Nodes_Pairs_1530

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Leaf struct {
	lv  int
	ind int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dist(lv1, ind1, lv2, ind2 int) int {
	cnt := 0
	if lv1 == lv2 {
		for ind2 > ind1 {
			ind2 /= 2
			ind1 /= 2
			cnt++
		}
		return 2 * cnt
	}
	for lv2 > lv1 {
		ind2 /= 2
		lv2--
		cnt++
	}
	return cnt + dist(lv1, min(ind1, ind2), lv2, max(ind1, ind2))
}

func postOrder(root *TreeNode, lv, ind int, leaf *[]Leaf) {
	if root == nil {
		return
	}
	postOrder(root.Left, lv+1, 2*ind, leaf)
	postOrder(root.Right, lv+1, 2*ind+1, leaf)
	if root.Left == nil && root.Right == nil {
		*leaf = append(*leaf, Leaf{lv, ind})
	}
}

func countPairs(root *TreeNode, d int) int {
	leaf := make([]Leaf, 0)
	postOrder(root, 1, 1, &leaf)

	sort.Slice(leaf, func(i, j int) bool {
		if leaf[i].lv == leaf[j].lv {
			return leaf[i].ind < leaf[j].ind
		}
		return leaf[i].lv < leaf[j].lv
	})
	ans := 0
	m := len(leaf)
	if m == 1 {
		return 0
	}
	for i := 0; i < m-1; i++ {
		cur := leaf[i]
		for j := i + 1; j < m; j++ {
			tmp := leaf[j]
			if tmp.lv-cur.lv > d {
				break
			}
			if dist(cur.lv, cur.ind, tmp.lv, tmp.ind) <= d {
				ans++
			}
		}
	}
	return ans
}
