// https://leetcode.com/problems/add-digits/
package main

func addDigits(num int) int {
	for tmp := 0; num >= 10; num, tmp = tmp, num {
		for ; num > 0; num /= 10 {
			tmp += num % 10
		}
	}
	return num
}

func main() {}
