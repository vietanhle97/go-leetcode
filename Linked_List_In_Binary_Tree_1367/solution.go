package Linked_List_In_Binary_Tree_1367

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func check(root *TreeNode, linkedList []int, res *bool) {
	if *res {
		return
	}
	if len(linkedList) == 0 {
		*res = true
		return
	}
	if root == nil {
		return
	}
	if root.Val == linkedList[0] {
		check(root.Left, linkedList[1:], res)
		check(root.Right, linkedList[1:], res)
	} else {
		return
	}
}

func preOrder(root *TreeNode, linkedList []int, res *bool) {
	if *res {
		return
	}
	if root == nil {
		return
	}
	check(root, linkedList, res)
	preOrder(root.Left, linkedList, res)
	preOrder(root.Right, linkedList, res)
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	res := false
	linkedList := make([]int, 0)
	for head != nil {
		linkedList = append(linkedList, head.Val)
		head = head.Next
	}
	preOrder(root, linkedList, &res)
	return res
}
