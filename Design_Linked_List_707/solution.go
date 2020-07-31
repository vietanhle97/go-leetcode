package Design_Linked_List_707

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
	size int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	head := &Node{-1, nil}
	return MyLinkedList{head, 0}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (list *MyLinkedList) Get(index int) int {
	if index >= list.size || list.IsEmpty() {
		return -1
	}
	cur, cnt := list.head.next, 0
	for cnt < index {
		cur = cur.next
		cnt++
	}
	return cur.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (list *MyLinkedList) AddAtHead(val int) {
	list.size++
	if list.IsEmpty() {
		list.head.next = &Node{val, nil}
		return
	}
	tmp, node := list.head.next, &Node{val, nil}
	list.head.next = node
	node.next = tmp
}

/** Append a node of value val to the last element of the linked list. */
func (list *MyLinkedList) AddAtTail(val int) {
	list.size++
	if list.IsEmpty() {
		list.head.next = &Node{val, nil}
		return
	}
	tmp := list.head
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = &Node{val, nil}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index > list.size {
		return
	}
	list.size++
	node := &Node{val, nil}
	prev, cur, cnt := list.head, list.head.next, 0
	for cnt < index {
		prev, cur = prev.next, cur.next
		cnt++
	}
	prev.next, node.next = node, cur
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (list *MyLinkedList) DeleteAtIndex(index int) {
	if list.IsEmpty() || index >= list.size {
		return
	}
	prev, cur, cnt := list.head, list.head.next, 0
	for cnt < index {
		prev, cur = prev.next, cur.next
		cnt++
	}
	prev.next = cur.next
	list.size--
}

func (list *MyLinkedList) IsEmpty() bool {
	return list.head.next == nil
}
