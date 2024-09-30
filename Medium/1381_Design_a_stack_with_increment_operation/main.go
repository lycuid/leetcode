// https://leetcode.com/problems/design-a-stack-with-increment-operation/
package main

import "container/list"

type CustomStack struct {
	capacity int
	items    *list.List
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		capacity: maxSize,
		items:    list.New(),
	}
}

func (this *CustomStack) Push(x int) {
	if this.items.Len() < this.capacity {
		this.items.PushBack(x)
	}
}

func (this *CustomStack) Pop() int {
	if this.items.Len() > 0 {
		return this.items.Remove(this.items.Back()).(int)
	}
	return -1
}

func (this *CustomStack) Increment(k int, val int) {
	node := this.items.Front()
	for i := 0; i < k && node != nil; i++ {
		node.Value = node.Value.(int) + val
		node = node.Next()
	}
}

func main() {}
