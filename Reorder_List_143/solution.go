package Reorder_List_143

type ListNode struct {
	Val  int
	Next *ListNode
}

var prev, cur *ListNode

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev, cur = nil, slow.Next
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	slow.Next = nil
	h1, h2 := head, prev
	for h2 != nil {
		nx := h1.Next
		h1.Next = h2
		h1 = h2
		h2 = nx
	}
}
