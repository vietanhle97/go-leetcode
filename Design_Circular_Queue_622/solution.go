package Design_Circular_Queue_622

type MyCircularQueue struct {
	data  []int
	front int
	rear  int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	data := make([]int, k+1)
	return MyCircularQueue{data, 0, 0}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (buffer *MyCircularQueue) EnQueue(value int) bool {
	if buffer.IsFull() {
		return false
	}
	buffer.data[buffer.rear] = value
	buffer.rear = (buffer.rear + 1) % (len(buffer.data))
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (buffer *MyCircularQueue) DeQueue() bool {
	if buffer.IsEmpty() {
		return false
	}
	buffer.data[buffer.front] = 0
	buffer.front = (buffer.front + 1) % (len(buffer.data))
	return true
}

/** Get the front item from the queue. */
func (buffer *MyCircularQueue) Front() int {
	if buffer.IsEmpty() {
		return -1
	}
	return buffer.data[buffer.front]
}

/** Get the last item from the queue. */
func (buffer *MyCircularQueue) Rear() int {
	if buffer.IsEmpty() {
		return -1
	}
	ind := buffer.rear - 1
	if ind < 0 {
		ind = len(buffer.data) + ind
	}
	return buffer.data[ind]
}

/** Checks whether the circular queue is empty or not. */
func (buffer *MyCircularQueue) IsEmpty() bool {
	return buffer.front == buffer.rear
}

/** Checks whether the circular queue is full or not. */
func (buffer *MyCircularQueue) IsFull() bool {
	return (buffer.rear+1)%(len(buffer.data)) == buffer.front
}
