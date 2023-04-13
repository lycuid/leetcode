// https://leetcode.com/problems/validate-stack-sequences/
package main

func validateStackSequences(pushed []int, popped []int) bool {
	for i, j := 0, 0; i < len(pushed); i, j = i+1, j+1 {
		pushed[j] = pushed[i]
		for j >= 0 && len(popped) > 0 && pushed[j] == popped[0] {
			j, popped = j-1, popped[1:]
		}
	}
	return len(popped) == 0
}

func main() {}
