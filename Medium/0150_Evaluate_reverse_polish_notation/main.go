// https://leetcode.com/problems/evaluate-reverse-polish-notation/
package main

import "strconv"

type Stack struct {
	inner  [1e4 + 1]int
	cursor int
}

func (this *Stack) Push(item int) {
	this.inner[this.cursor] = item
	this.cursor++
}

func (this *Stack) Pop() int {
	this.cursor--
	return this.inner[this.cursor]
}

func evalRPN(tokens []string) int {
	stack := Stack{}
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			b, a := stack.Pop(), stack.Pop()
			switch token {
			case "+":
				stack.Push(a + b)
			case "-":
				stack.Push(a - b)
			case "*":
				stack.Push(a * b)
			case "/":
				stack.Push(a / b)
			}
		default:
			num, _ := strconv.Atoi(token)
			stack.Push(num)
		}
	}
	return stack.Pop()
}

func main() {}
