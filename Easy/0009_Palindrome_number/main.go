// https://leetcode.com/problems/palindrome-number/

package main

func reverseNumber(num int) int {
	rnum := 0
	for num > 0 {
		rnum = rnum*10 + (num % 10)
		num /= 10
	}
	return rnum
}

func isPalindrome(num int) bool {
	reverse := reverseNumber(num)
	if reverse == num {
		return true
	}
	return false
}

func main() {}
