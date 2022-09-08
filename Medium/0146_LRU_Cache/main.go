// https://leetcode.com/problems/lru-cache/
package main

type List struct {
	key, value int
	prev, next *List
}

type LRUCache struct {
	c_head, c_end *List
	pairs         map[int]*List
	capacity      int
}

func (cache *LRUCache) Attach(l *List) {
	if l == nil {
		return
	}
	cache.capacity--
	// first element.
	if cache.c_head == nil {
		cache.c_end = l
	} else {
		cache.c_head.prev = l
	}
	l.next = cache.c_head
	l.prev = nil
	cache.c_head = l
}

func (cache *LRUCache) Detach(l *List) {
	if l == nil {
		return
	}
	cache.capacity++
	if l.prev != nil {
		l.prev.next = l.next
	} else {
		cache.c_head = l.next
	}
	if l.next != nil {
		l.next.prev = l.prev
	} else {
		cache.c_end = l.prev
	}
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		pairs:    make(map[int]*List),
		capacity: capacity,
	}
}

func (cache *LRUCache) Get(key int) int {
	if list, ok := cache.pairs[key]; ok {
		cache.Detach(list)
		cache.Attach(list)
		return list.value
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	// check if key exists, replace if exists and return.
	if list, ok := cache.pairs[key]; ok {
		list.value = value
		cache.Detach(list)
		cache.Attach(list)
		return
	}

	// create new list.
	l := List{key: key, value: value}
	cache.pairs[key] = &l
	cache.Attach(&l)
	if cache.capacity >= 0 {
		return
	}

	// pop least recently used.
	end := cache.c_end
	cache.Detach(end)
	cache.c_end = end.prev
	delete(cache.pairs, end.key)
}

func main() {}
