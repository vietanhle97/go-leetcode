package Rotate_List_61

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	tmp, l, cnt := head, 1, 1
	for tmp.Next != nil {
		tmp = tmp.Next
		l++
	}
	tmp.Next = head
	for cnt < l-(k%l) {
		head = head.Next
		cnt += 1
	}
	newHead := head.Next
	head.Next = nil
	return newHead
}
