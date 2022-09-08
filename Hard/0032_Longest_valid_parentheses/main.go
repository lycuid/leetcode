// https://leetcode.com/problems/longest-valid-parentheses/
package main

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Stack struct {
	inner  []int
	cursor int
}

func MakeStack(n int) Stack {
	return Stack{inner: make([]int, n)}
}

func (this *Stack) Push(item int) {
	this.inner[this.cursor] = item
	this.cursor++
}

func (this *Stack) Pop() int {
	this.cursor--
	return this.inner[this.cursor]
}

func (this *Stack) Top() int {
	return this.inner[this.cursor-1]
}

func (this *Stack) Empty() bool {
	return this.cursor == 0
}

func longestValidParentheses(str string) (count int) {
	stack := MakeStack(len(str)+1)
	stack.Push(-1)
	for i := range str {
		ch := str[i]
		if ch == '(' {
			stack.Push(i)
		} else {
			stack.Pop()
			if stack.Empty() {
				stack.Push(i)
			} else {
				count = Max(count, i-stack.Top())
			}
		}
	}
	return count
}

func main() {}
