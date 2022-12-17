// https://leetcode.com/problems/evaluate-reverse-polish-notation/
package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	Pop := func() (item int) {
		item, stack = stack[len(stack)-1], stack[:len(stack)-1]
		return item
	}
	for _, token := range tokens {
		switch token {
		case "+":
			snd, fst := Pop(), Pop()
			stack = append(stack, fst+snd)
		case "-":
			snd, fst := Pop(), Pop()
			stack = append(stack, fst-snd)
		case "*":
			snd, fst := Pop(), Pop()
			stack = append(stack, fst*snd)
		case "/":
			snd, fst := Pop(), Pop()
			stack = append(stack, fst/snd)
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

func main() {}
