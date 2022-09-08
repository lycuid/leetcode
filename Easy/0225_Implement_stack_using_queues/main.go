// https://leetcode.com/problems/implement-stack-using-queues/
package main

type MyStack struct {
	stack  []int
	cursor int
}

func Constructor() MyStack {
	return MyStack{stack: make([]int, 100)}
}

func (this *MyStack) Push(x int) {
	this.stack[this.cursor] = x
	this.cursor++
}

func (this *MyStack) Pop() int {
	this.cursor--
	return this.stack[this.cursor]
}

func (this *MyStack) Top() int {
	return this.stack[this.cursor-1]
}

func (this *MyStack) Empty() bool {
	return this.cursor == 0
}

func main() {}
