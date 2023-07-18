// https://leetcode.com/problems/lru-cache/
package main

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity    int
	inner       map[int]*Node
	head, stale *Node
}

func (cache *LRUCache) Attach(node *Node) {
	if node != nil {
		if node.next = cache.head; cache.head != nil {
			cache.head.prev = node
		} else {
			cache.stale = node
		}
		cache.head, cache.capacity = node, cache.capacity-1
	}
}

func (cache *LRUCache) Detach(node *Node) *Node {
	if node != nil {
		if node.prev != nil {
			node.prev.next = node.next
		} else {
			cache.head = node.next
		}
		if node.next != nil {
			node.next.prev = node.prev
		} else {
			cache.stale = node.prev
		}
		node.next, node.prev, cache.capacity = nil, nil, cache.capacity+1
	}
	return node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, inner: make(map[int]*Node)}
}

func (cache *LRUCache) Get(key int) int {
	if node := cache.inner[key]; node != nil {
		cache.Attach(cache.Detach(node))
		return node.value
	}
	return -1
}

func (cache *LRUCache) Put(key, value int) {
	node := cache.Detach(cache.inner[key])
	if node == nil {
		if cache.capacity == 0 {
			if stale := cache.Detach(cache.stale); stale != nil {
				delete(cache.inner, stale.key)
			}
		}
		node = &Node{key: key, value: value}
		cache.inner[key] = node
	}
	node.value = value
	cache.Attach(node)
}

func main() {}
