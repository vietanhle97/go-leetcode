package Add_Two_Numbers_2

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carrier := 0
	res := &ListNode{-1, nil}
	tmp := res
	for l1 != nil || l2 != nil {
		val1 := 0
		val2 := 0
		if tmp.Val == -1 {
			if l1 != nil {
				val1 = l1.Val
			}
			if l2 != nil {
				val2 = l2.Val
			}
			tmp.Val = (val1 + val2 + carrier) % 10
			carrier = (val1 + val2 + carrier) / 10
		} else {
			if l1 != nil {
				val1 = l1.Val
			}
			if l2 != nil {
				val2 = l2.Val
			}
			tmp.Next = &ListNode{(val1 + val2 + carrier) % 10, nil}
			carrier = (val1 + val2 + carrier) / 10
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if tmp.Next != nil {
			tmp = tmp.Next
		}
	}
	if carrier == 1 {
		tmp.Next = &ListNode{1, nil}
	}
	return res
}
