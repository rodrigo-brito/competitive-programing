package main

type LRUCache struct {
	Capacity  int
	Items     map[int]int
	OrderKeys []int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity:  capacity,
		OrderKeys: make([]int, 0, capacity),
		Items:     make(map[int]int),
	}
}

func (l *LRUCache) Get(key int) int {
	value, found := l.Items[key]
	if !found {
		return -1
	}

	for i, k := range l.OrderKeys {
		if k == key {
			l.OrderKeys = append(l.OrderKeys[:i], l.OrderKeys[i+1:]...)
			break
		}
	}
	l.OrderKeys = append(l.OrderKeys, key)
	return value
}

func (l *LRUCache) Put(key int, value int) {
	_, found := l.Items[key]
	if found {
		for i, k := range l.OrderKeys {
			if k == key {
				l.OrderKeys = append(l.OrderKeys[:i], l.OrderKeys[i+1:]...)
				break
			}
		}
		l.OrderKeys = append(l.OrderKeys, key)
	} else {
		if len(l.OrderKeys) == l.Capacity {
			delete(l.Items, l.OrderKeys[0])
			l.OrderKeys = l.OrderKeys[1:]
		}

		l.OrderKeys = append(l.OrderKeys, key)
	}
	l.Items[key] = value
}
