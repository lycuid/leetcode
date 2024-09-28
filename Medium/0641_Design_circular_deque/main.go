// https://leetcode.com/problems/design-circular-deque/
package main

import "container/list"

type MyCircularDeque struct {
	items *list.List
	size  int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{list.New(), k}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.size == this.items.Len() {
		return false
	}
	this.items.PushFront(value)
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.size == this.items.Len() {
		return false
	}
	this.items.PushBack(value)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.items.Remove(this.items.Front())
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.items.Remove(this.items.Back())
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if item := this.items.Front(); item != nil {
		return item.Value.(int)
	}
	return -1
}

func (this *MyCircularDeque) GetRear() int {
	if item := this.items.Back(); item != nil {
		return item.Value.(int)
	}
	return -1
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.items.Len() == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size == this.items.Len()
}

func main() {}
