// https://leetcode.com/problems/implement-queue-using-stacks/
package main

type Stack []int

func (stack Stack) Empty() bool    { return len(stack) == 0 }
func (stack Stack) Top() int       { return stack[len(stack)-1] }
func (stack *Stack) Push(item int) { *stack = append((*stack), item) }
func (stack *Stack) Pop() (item int) {
	item, *stack = (*stack)[len(*stack)-1], (*stack)[:len(*stack)-1]
	return item
}

type MyQueue struct{ in, out Stack }

func Constructor() MyQueue { return MyQueue{} }

func (this *MyQueue) Push(x int) { this.in.Push(x) }

func (this *MyQueue) adjust() {
	if this.out.Empty() {
		for !this.in.Empty() {
			this.out.Push(this.in.Pop())
		}
	}
}

func (this *MyQueue) Pop() int {
	this.adjust()
	return this.out.Pop()
}

func (this *MyQueue) Peek() int {
	this.adjust()
	return this.out.Top()
}

func (this *MyQueue) Empty() bool {
	return this.in.Empty() && this.out.Empty()
}

func main() {}
