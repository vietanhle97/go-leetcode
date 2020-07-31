package LRU_Cache_146

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (dll *DoublyLinkedList) insert(node *Node) {
	tmp := dll.head.next
	dll.head.next, node.next = node, tmp
	tmp.prev, node.prev = node, dll.head
	dll.size++
}

func (dll *DoublyLinkedList) remove(node *Node) {
	prev, next := node.prev, node.next
	prev.next, next.prev = next, prev
	node.next, node.prev = nil, nil
	dll.size--
}

func (dll *DoublyLinkedList) removeTail() *Node {
	node := dll.tail.prev
	dll.remove(node)
	return node
}
func (dll *DoublyLinkedList) move(node *Node) {
	if node.next != nil && node.prev != nil {
		dll.remove(node)
	}
	dll.insert(node)
}

type LRUCache struct {
	dll      *DoublyLinkedList
	num      map[int]*Node
	capacity int
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.next, tail.next = tail, head
	head.prev, tail.prev = tail, head
	dll := &DoublyLinkedList{head, tail, 0}
	return LRUCache{dll, map[int]*Node{}, capacity}
}

func (lru *LRUCache) Get(key int) int {
	if _, ok := lru.num[key]; !ok {
		return -1
	}
	node := lru.num[key]
	lru.dll.move(node)
	return node.val
}

func (lru *LRUCache) Put(key int, value int) {
	node := &Node{key: key}
	if _, ok := lru.num[key]; ok {
		node = lru.num[key]
	}
	node.val = value
	lru.dll.move(node)
	lru.num[key] = node
	if lru.dll.size > lru.capacity {
		delete(lru.num, lru.dll.removeTail().key)
	}
}
