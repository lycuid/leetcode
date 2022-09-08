// https://leetcode.com/problems/valid-parenthesis-string/
package main

func checkValidString(s string) bool {
	var any, i, bal int
	for _, ch := range s {
		if ch == '*' {
			any++
		}
	}
	for i = 0; i < len(s); i++ {
		if s[i] == ')' {
			bal--
		} else {
			bal++
		}
		if bal < 0 {
			return false
		}
	}
	for i, bal = len(s)-1, 0; i >= 0; i-- {
		if s[i] == '(' {
			bal--
		} else {
			bal++
		}
		if bal < 0 {
			return false
		}
	}
	return true
}

func main() {}
