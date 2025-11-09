// https://leetcode.com/problems/count-operations-to-obtain-zero/
package main

func countOperations(num1 int, num2 int) int {
	if num1 < num2 {
		num1, num2 = num2, num1
	}
	if num2 == 0 {
		return 0
	}
	return num1/num2 + countOperations(num2, num1%num2)
}

func main() {}
