// https://leetcode.com/problems/plus-one/
package main

func plusOne(digits []int) []int {
	for i := 0; i < len(digits)/2; i++ {
		digits[i], digits[len(digits)-1-i] = digits[len(digits)-1-i], digits[i]
	}
	for carry, i := 1, 0; carry > 0; i++ {
		if i >= len(digits) {
			digits = append(digits, 0)
		}
		carry += digits[i]
		digits[i], carry = carry%10, carry/10
	}
	for i := 0; i < len(digits)/2; i++ {
		digits[i], digits[len(digits)-1-i] = digits[len(digits)-1-i], digits[i]
	}
	return digits
}

func main() {}
