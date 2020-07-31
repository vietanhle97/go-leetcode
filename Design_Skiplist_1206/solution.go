package Design_Skiplist_1206

import "math/rand"

type Node struct {
	key     int
	forward []*Node
}

type Skiplist struct {
	MAX   int
	level int
	p     float64
	head  *Node
}

func (list *Skiplist) randomLevel() int {
	lvl := 0
	tmp := rand.Float64()
	for tmp < list.p && lvl < list.MAX {
		lvl += 1
		tmp = rand.Float64()
	}
	return lvl
}

func Constructor() Skiplist {
	MAX, level, p := 16, 0, 0.5
	forward := make([]*Node, MAX+1)
	head := &Node{-1, forward}
	return Skiplist{MAX, level, p, head}
}

func (list *Skiplist) Search(target int) bool {
	cur := list.head
	for i := list.level; i > -1; i-- {
		for cur.forward[i] != nil && cur.forward[i].key < target {
			cur = cur.forward[i]
		}
	}
	cur = cur.forward[0]
	return cur != nil && cur.key == target
}

func (list *Skiplist) Add(key int) {
	update, cur := make([]*Node, list.MAX+1), list.head
	for i := list.level; i > -1; i-- {
		for cur.forward[i] != nil && cur.forward[i].key <= key {
			cur = cur.forward[i]
		}
		update[i] = cur
	}
	cur = cur.forward[0]
	if cur == nil || cur.key != key {
		rlevel := list.randomLevel()
		if rlevel > list.level {
			for i := list.level + 1; i < rlevel+1; i++ {
				update[i] = list.head
			}
			list.level = rlevel
		}
		n := &Node{key, make([]*Node, rlevel+1)}
		for i := 0; i < rlevel+1; i++ {
			n.forward[i] = update[i].forward[i]
			update[i].forward[i] = n
		}
	}
}

func (list *Skiplist) Erase(key int) bool {
	if !list.Search(key) {
		return false
	}

	update, cur := make([]*Node, list.MAX+1), list.head
	for i := list.level; i > -1; i-- {
		for cur.forward[i] != nil && cur.forward[i].key < key {
			cur = cur.forward[i]
		}
		update[i] = cur
	}
	cur = cur.forward[0]
	if cur != nil && cur.key == key {
		for i := 0; i < list.level+1; i++ {
			if update[i].forward[i] != cur {
				break
			}
			update[i].forward[i] = cur.forward[i]
		}
		for list.level > 0 && list.head.forward[list.level] == nil {
			list.level--
		}
	}
	return true
}
