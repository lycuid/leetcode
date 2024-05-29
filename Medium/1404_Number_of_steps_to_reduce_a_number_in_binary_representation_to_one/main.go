// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/
package main

func numSteps(s string) (steps int) {
	i, j := len(s)-1, 0
	for ; i >= 0 && s[i] == '0'; i-- {
		steps++
	}
	for ; i > 0; i = j {
		for j = i - 1; j >= 0 && s[j] == '1'; j-- {
		}
		steps += i - j + 1
	}
	return steps
}

func main() {}
