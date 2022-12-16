// https://leetcode.com/problems/implement-queue-using-stacks/
package main

type Stack struct{ inner []int }

func (this *Stack) Push(item int) { this.inner = append(this.inner, item) }
func (this *Stack) Peek() int     { return this.inner[len(this.inner)-1] }
func (this *Stack) Empty() bool   { return len(this.inner) == 0 }
func (this *Stack) Pop() (item int) {
	item, this.inner = this.inner[len(this.inner)-1], this.inner[:len(this.inner)-1]
	return item
}

type MyQueue struct{ inner, alt Stack }

func Constructor() MyQueue { return MyQueue{} }
func (this *MyQueue) Push(item int) {
	for !this.inner.Empty() {
		this.alt.Push(this.inner.Pop())
	}
	this.inner.Push(item)
	for !this.alt.Empty() {
		this.inner.Push(this.alt.Pop())
	}
}
func (this *MyQueue) Pop() int    { return this.inner.Pop() }
func (this *MyQueue) Peek() int   { return this.inner.Peek() }
func (this *MyQueue) Empty() bool { return this.inner.Empty() }

func main() {}
