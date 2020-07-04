package Linked_List_Cycle_II_142

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	m := map[*ListNode]int{}
	cnt := 0
	for head.Next != nil {
		m[head] = cnt
		cnt++
		if head.Next == nil {
			return nil
		}
		if _, ok := m[head.Next]; ok {
			return head.Next
		}
		head = head.Next
	}
	return nil
}
