package main

type Node struct {
	Key        int
	Value      int
	next, prev *Node
}

type List struct {
	size       int
	head, tail *Node
}

func (l *List) Push(node *Node) {
	l.size++
	if l.head == nil {
		l.head = node // First node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
}

func (l *List) Pop() {
	l.size--
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	}
}

func (l *List) Delete(pointer *Node) {
	l.size--
	prev := pointer.prev
	next := pointer.next
	if pointer == l.head {
		l.head = pointer.next
	}
	if pointer == l.tail {
		l.tail = pointer.prev
	}
	if prev != nil {
		prev.next = pointer.next
	}
	if next != nil {
		next.prev = pointer.prev
	}
	pointer.prev = nil
	pointer.next = nil
}

type LRUCache struct {
	Capacity int
	Indexes  map[int]*Node
	Items    *List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Items:    new(List),
		Indexes:  make(map[int]*Node),
	}
}

func (l *LRUCache) Get(key int) int {
	pointer, found := l.Indexes[key]
	if !found {
		return -1
	}

	if pointer != l.Items.tail {
		l.Items.Delete(pointer)
		l.Items.Push(pointer)
	}

	return pointer.Value
}

func (l *LRUCache) Put(key int, value int) {
	pointer, found := l.Indexes[key]
	if found {
		l.Items.Delete(pointer)
		pointer.Value = value
		l.Items.Push(pointer)
	} else {
		if l.Items.size == l.Capacity {
			delete(l.Indexes, l.Items.head.Key)
			l.Items.Pop()
		}

		node := &Node{Key: key, Value: value}
		l.Items.Push(node)
		l.Indexes[key] = node
	}
}
