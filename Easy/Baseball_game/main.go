// https://leetcode.com/problems/baseball-game/
package main

import "strconv"

const MAX_OPS = 1001

func calPoints(ops []string) (sum int) {
	stack := make([]int, MAX_OPS)
	cursor := 0
	for _, op := range ops {
		switch op {
		case "+":
			stack[cursor] = stack[cursor-1] + stack[cursor-2]
			break
		case "D":
			stack[cursor] = stack[cursor-1] * 2
			break
		case "C":
			cursor -= 2
			break
		default:
			i, err := strconv.Atoi(op)
			if err != nil {
				panic("I find the lack of integer, Disturbing!.")
			}
			stack[cursor] = i
			break
		}
		cursor++
	}
	for i := 0; i < cursor; i++ {
		sum += stack[i]
	}
	return sum
}

func main() {}
