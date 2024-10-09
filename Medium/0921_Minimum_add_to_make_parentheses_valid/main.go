// https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/
package main

func minAddToMakeValid(s string) (count int) {
	var open int
	for _, ch := range s {
		switch ch {
		case '(':
			open++
		case ')':
			if open > 0 {
				open--
			} else {
				count++
			}
		}
	}
	return count + open
}

func main() {}
