// https://leetcode.com/problems/design-hashset/
package main

type Node struct {
	Key  int
	Next *Node
}

type MyHashSet struct{ inner []*Node }

func Constructor() MyHashSet {
	return MyHashSet{inner: make([]*Node, 100)}
}

func (this MyHashSet) Hash(key int) int {
	return key % len(this.inner)
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		this.inner[this.Hash(key)] = &Node{key, this.inner[this.Hash(key)]}
	}
}

func (this *MyHashSet) Remove(key int) {
	var (
		prev *Node = nil
		curr *Node = this.inner[this.Hash(key)]
	)
	for curr != nil {
		if curr.Key == key {
			if prev == nil {
				this.inner[this.Hash(key)] = curr.Next
			} else {
				prev.Next = curr.Next
			}
			curr.Next = nil
		}
		prev, curr = curr, curr.Next
	}
}

func (this *MyHashSet) Contains(key int) bool {
	for node := this.inner[this.Hash(key)]; node != nil; node = node.Next {
		if node.Key == key {
			return true
		}
	}
	return false
}

func main() {}
