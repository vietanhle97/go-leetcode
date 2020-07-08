package Smallest_String_Starting_From_Leaf_988

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b string) string {
	if a > b {
		return b
	}
	return a
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func preOrder(root *TreeNode, cur string, res *string, m map[int]string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if len(*res) == 0 {
			*res = m[root.Val] + reverse(cur)
		} else {
			*res = min(*res, m[root.Val]+reverse(cur))
		}
	}
	preOrder(root.Left, cur+m[root.Val], res, m)
	preOrder(root.Right, cur+m[root.Val], res, m)

}

func smallestFromLeaf(root *TreeNode) string {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	m := map[int]string{}
	for i := 0; i < 26; i++ {
		m[i] = string(alpha[i])
	}
	res := ""
	preOrder(root, "", &res, m)
	return res
}
