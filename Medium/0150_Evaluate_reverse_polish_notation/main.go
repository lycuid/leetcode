// https://leetcode.com/problems/evaluate-reverse-polish-notation/
package main

import "strconv"

type Stack []int

func (stack *Stack) Push(item int) { *stack = append(*stack, item) }

func (stack *Stack) Pop() (item int) {
	item, *stack = (*stack)[len(*stack)-1], (*stack)[:len(*stack)-1]
	return item
}

func evalRPN(tokens []string) int {
	var stack Stack
	for _, token := range tokens {
		switch token {
		case "+":
			stack.Push(stack.Pop() + stack.Pop())
			break
		case "-":
			rhs, lhs := stack.Pop(), stack.Pop()
			stack.Push(lhs - rhs)
			break
		case "*":
			stack.Push(stack.Pop() * stack.Pop())
			break
		case "/":
			rhs, lhs := stack.Pop(), stack.Pop()
			stack.Push(lhs / rhs)
			break
		default:
			if num, err := strconv.Atoi(token); err == nil {
				stack = append(stack, num)
			}
		}
	}
	return stack[0]
}

func main() {}
