// https://leetcode.com/problems/valid-parentheses/
package main

type Stack struct {
	inner  []rune
	cursor int
}

func (this *Stack) Push(item rune) {
	this.inner[this.cursor] = item
	this.cursor++
}

func (this *Stack) Pop() (item rune) {
	this.cursor--
	item = this.inner[this.cursor]
	return item
}

func (this *Stack) Empty() bool {
	return this.cursor == 0
}

func isValid(str string) bool {
	stack := Stack{make([]rune, len(str)), 0}
	for _, ch := range str {
		switch ch {
		case 40, 91, 123:
			stack.Push(ch)
		case 41:
			if stack.Empty() || stack.Pop() != 40 {
				return false
			}
		case 93:
			if stack.Empty() || stack.Pop() != 91 {
				return false
			}
		case 125:
			if stack.Empty() || stack.Pop() != 123 {
				return false
			}
		}
	}
	return stack.Empty()
}

func main() {}
