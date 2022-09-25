// https://leetcode.com/problems/design-circular-queue/
package main

type MyCircularQueue struct {
	inner       []int
	front, rear int
}

const EMPTY = -1

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{make([]int, k), EMPTY, EMPTY}
}

func (this *MyCircularQueue) EnQueue(value int) (cond bool) {
	if cond = !this.IsFull(); cond {
		if this.IsEmpty() {
			this.front = 0
		}
		this.rear = (this.rear + 1) % len(this.inner)
		this.inner[this.rear] = value
	}
	return cond
}

func (this *MyCircularQueue) DeQueue() (cond bool) {
	if cond = !this.IsEmpty(); cond {
		if this.rear == this.front {
			this.front, this.rear = EMPTY, EMPTY
		} else {
			this.front = (this.front + 1) % len(this.inner)
		}
	}
	return cond
}

func (this MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.inner[this.front]
}

func (this MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.inner[this.rear]
}

func (this MyCircularQueue) IsEmpty() bool {
	return this.front == EMPTY
}

func (this MyCircularQueue) IsFull() bool {
	return this.front == (this.rear+1)%len(this.inner)
}

func main() {}
