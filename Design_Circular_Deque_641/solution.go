package Design_Circular_Deque_641

type Node struct {
	val  int
	next *Node
	prev *Node
}

type MyCircularDeque struct {
	front *Node
	rear  *Node
	cnt   int
	size  int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	front, rear := &Node{}, &Node{}
	front.next, rear.next = rear, front
	front.prev, rear.prev = rear, front
	return MyCircularDeque{front, rear, 0, k}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (deque *MyCircularDeque) InsertFront(value int) bool {
	if deque.IsFull() {
		return false
	}
	node, tmp := &Node{value, nil, nil}, deque.front.next
	deque.front.next, node.next = node, tmp
	node.prev, tmp.prev = deque.front, node
	deque.cnt++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (deque *MyCircularDeque) InsertLast(value int) bool {
	if deque.IsFull() {
		return false
	}
	node, tmp := &Node{value, nil, nil}, deque.rear.prev
	node.next, tmp.next = deque.rear, node
	deque.rear.prev, node.prev = node, tmp
	deque.cnt++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (deque *MyCircularDeque) DeleteFront() bool {
	if deque.IsEmpty() {
		return false
	}
	tmp := deque.front.next
	deque.front.next = tmp.next
	tmp.next.prev = deque.front
	deque.cnt--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (deque *MyCircularDeque) DeleteLast() bool {
	if deque.IsEmpty() {
		return false
	}
	tmp := deque.rear.prev
	deque.rear.prev = tmp.prev
	tmp.prev.next = deque.rear
	deque.cnt--
	return true
}

/** Get the front item from the deque. */
func (deque *MyCircularDeque) GetFront() int {
	if deque.IsEmpty() {
		return -1
	}
	return deque.front.next.val
}

/** Get the last item from the deque. */
func (deque *MyCircularDeque) GetRear() int {
	if deque.IsEmpty() {
		return -1
	}
	return deque.rear.prev.val
}

/** Checks whether the circular deque is empty or not. */
func (deque *MyCircularDeque) IsEmpty() bool {
	return deque.cnt == 0
}

/** Checks whether the circular deque is full or not. */
func (deque *MyCircularDeque) IsFull() bool {
	return deque.cnt == deque.size
}
