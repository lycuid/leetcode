// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/
package main

func numberOfSteps(num int) (steps int) {
	for num > 0 {
		if num%2 == 0 {
			num >>= 1
		} else {
			num -= 1
		}
		steps++
	}
	return steps
}

func main() {}
